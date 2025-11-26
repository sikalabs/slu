package postgres_utils

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"time"
)

func RunPSQL(ctx context.Context, host string, port int, user, password, db, sql string) error {
	args := []string{
		"-h", host,
		"-p", fmt.Sprint(port),
		"-U", user,
		"-d", db,
		"-v", "ON_ERROR_STOP=1",
		"-c", sql,
	}
	cmd := exec.CommandContext(ctx, "psql", args...)
	cmd.Env = append(os.Environ(), "PGPASSWORD="+password)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func RunPSQLCapture(ctx context.Context, host string, port int, user, password, db, sql string, out *strings.Builder) error {
	args := []string{
		"-h", host,
		"-p", fmt.Sprint(port),
		"-U", user,
		"-d", db,
		"-v", "ON_ERROR_STOP=1",
		"-At",
		"-c", sql,
	}
	cmd := exec.CommandContext(ctx, "psql", args...)
	cmd.Env = append(os.Environ(), "PGPASSWORD="+password)
	cmd.Stdout = out
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func RunPSQLFromReader(ctx context.Context, host string, port int, user, password, db string, r io.Reader) error {
	args := []string{
		"-h", host,
		"-p", fmt.Sprint(port),
		"-U", user,
		"-d", db,
		"-v", "ON_ERROR_STOP=0",
	}

	cmd := exec.CommandContext(ctx, "psql", args...)
	cmd.Env = append(os.Environ(), "PGPASSWORD="+password)
	cmd.Stdin = r

	// Capture stdout + stderr
	var stderr bytes.Buffer
	cmd.Stdout = os.Stdout
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err == nil {
		return nil
	}

	// ✅ Ignore known harmless pg_ errors
	out := stderr.String()
	if strings.Contains(out, "pg_replication_origin_advance") ||
		strings.Contains(out, "pg_catalog") ||
		strings.Contains(out, "pg_stat_") ||
		strings.Contains(out, "pg_trgm") ||
		strings.Contains(out, "pglogical") {
		fmt.Println("⚠️  Ignoring non-critical pg_ error(s) during restore.")
		fmt.Println(strings.TrimSpace(out))
		return nil
	}

	fmt.Fprint(os.Stderr, out)
	return fmt.Errorf("psql failed: %w", err)
}

func RestrictPSQLConnections(ctx context.Context, host string, port int, adminUser, adminPass, adminDB, targetDB, allowedUser string) error {
	if err := RunPSQL(ctx, host, port, adminUser, adminPass, adminDB,
		fmt.Sprintf(`REVOKE CONNECT ON DATABASE %q FROM PUBLIC;`, targetDB)); err != nil {
		return err
	}

	revokeExplicit := fmt.Sprintf(`
DO $$
DECLARE
  dbname       text := %s;
  grantee_oid  oid;
  grantee_name text;
BEGIN
  FOR grantee_oid IN
    WITH acl AS (
      SELECT (aclexplode(datacl)).*
      FROM pg_database
      WHERE datname = dbname
    )
    SELECT grantee
    FROM acl
    WHERE privilege_type = 'CONNECT'
  LOOP
    grantee_name := pg_get_userbyid(grantee_oid);
    IF grantee_name IS NOT NULL AND grantee_name NOT IN (%s, 'postgres') THEN
      EXECUTE 'REVOKE CONNECT ON DATABASE '
             || quote_ident(dbname)
             || ' FROM '
             || quote_ident(grantee_name);
    END IF;
  END LOOP;
END$$;`,
		QuoteLiteral(targetDB),
		QuoteLiteral(allowedUser),
	)
	if err := RunPSQL(ctx, host, port, adminUser, adminPass, adminDB, revokeExplicit); err != nil {
		return err
	}

	if err := RunPSQL(ctx, host, port, adminUser, adminPass, adminDB,
		fmt.Sprintf(`GRANT CONNECT ON DATABASE %q TO %q, postgres;`, targetDB, allowedUser)); err != nil {
		return err
	}

	killSQL := fmt.Sprintf(`
SELECT pg_terminate_backend(pid)
FROM pg_stat_activity
WHERE datname = %s
  AND usename <> %s
  AND pid <> pg_backend_pid();`,
		QuoteLiteral(targetDB),
		QuoteLiteral(allowedUser),
	)

	checkSQL := fmt.Sprintf(`
SELECT COUNT(*)
FROM pg_stat_activity
WHERE datname = %s
  AND usename <> %s
  AND pid <> pg_backend_pid();`,
		QuoteLiteral(targetDB),
		QuoteLiteral(allowedUser),
	)

	for i := 0; i < 25; i++ { // ~5s max
		if err := RunPSQL(ctx, host, port, adminUser, adminPass, adminDB, killSQL); err != nil {
			return err
		}
		var out strings.Builder
		if err := RunPSQLCapture(ctx, host, port, adminUser, adminPass, adminDB, checkSQL, &out); err != nil {
			return err
		}
		if strings.TrimSpace(out.String()) == "0" {
			break
		}
		time.Sleep(200 * time.Millisecond)
	}

	return nil
}

func ReopenPSQLConnections(ctx context.Context, host string, port int, adminUser, adminPass, adminDB, targetDB string) error {
	sql := fmt.Sprintf(`GRANT CONNECT ON DATABASE "%s" TO PUBLIC;`, targetDB)
	return RunPSQL(ctx, host, port, adminUser, adminPass, adminDB, sql)
}

func PurgePSQLPublicSchemaObjects(ctx context.Context, host string, port int, user, password, db string) error {
	// 1) Drop materialized views
	// Skip extension-owned materialized views
	dropMatViews := `
DO $$
DECLARE r record;
BEGIN
  FOR r IN
    SELECT c.oid, n.nspname AS schemaname, c.relname AS matviewname
    FROM pg_class c
    JOIN pg_namespace n ON n.oid = c.relnamespace
    WHERE n.nspname = 'public'
      AND c.relkind = 'm'
      AND NOT EXISTS (
        SELECT 1 FROM pg_depend d
        WHERE d.objid = c.oid
          AND d.deptype = 'e'
      )
  LOOP
    EXECUTE format('DROP MATERIALIZED VIEW IF EXISTS %I.%I CASCADE', r.schemaname, r.matviewname);
  END LOOP;
END$$;`
	if err := RunPSQL(ctx, host, port, user, password, db, dropMatViews); err != nil {
		return err
	}

	// 2) Drop normal views
	// Skip extension-owned views to preserve extensions like pg_stat_statements
	dropViews := `
DO $$
DECLARE r record;
BEGIN
  FOR r IN
    SELECT c.oid, n.nspname AS schemaname, c.relname AS viewname
    FROM pg_class c
    JOIN pg_namespace n ON n.oid = c.relnamespace
    WHERE n.nspname = 'public'
      AND c.relkind = 'v'
      AND NOT EXISTS (
        SELECT 1 FROM pg_depend d
        WHERE d.objid = c.oid
          AND d.deptype = 'e'
      )
  LOOP
    EXECUTE format('DROP VIEW IF EXISTS %I.%I CASCADE', r.schemaname, r.viewname);
  END LOOP;
END$$;`
	if err := RunPSQL(ctx, host, port, user, password, db, dropViews); err != nil {
		return err
	}

	// 3) Drop FOREIGN TABLES  (FIXED: use information_schema column names)
	dropForeignTables := `
DO $$
DECLARE r record;
BEGIN
  FOR r IN
    SELECT foreign_table_schema AS schemaname,
           foreign_table_name   AS relname
    FROM information_schema.foreign_tables
    WHERE foreign_table_schema = 'public'
  LOOP
    EXECUTE format('DROP FOREIGN TABLE IF EXISTS %I.%I CASCADE', r.schemaname, r.relname);
  END LOOP;
END$$;`
	if err := RunPSQL(ctx, host, port, user, password, db, dropForeignTables); err != nil {
		return err
	}

	// 4) Drop ordinary tables (indexes/constraints/triggers go with CASCADE)
	dropTables := `
DO $$
DECLARE r record;
BEGIN
  FOR r IN
    SELECT schemaname, tablename
    FROM pg_tables
    WHERE schemaname = 'public'
  LOOP
    EXECUTE format('DROP TABLE IF EXISTS %I.%I CASCADE', r.schemaname, r.tablename);
  END LOOP;
END$$;`
	if err := RunPSQL(ctx, host, port, user, password, db, dropTables); err != nil {
		return err
	}

	// 5) Drop standalone sequences
	dropSequences := `
DO $$
DECLARE r record;
BEGIN
  FOR r IN
    SELECT sequence_schema, sequence_name
    FROM information_schema.sequences
    WHERE sequence_schema = 'public'
  LOOP
    EXECUTE format('DROP SEQUENCE IF EXISTS %I.%I CASCADE', r.sequence_schema, r.sequence_name);
  END LOOP;
END$$;`
	if err := RunPSQL(ctx, host, port, user, password, db, dropSequences); err != nil {
		return err
	}

	// 6) Drop routines (functions/procedures/aggregates) by identity signature
	// Skip extension-owned functions to preserve extensions like pg_stat_statements, pg_trgm
	dropRoutines := `
DO $$
DECLARE r record;
BEGIN
  FOR r IN
    SELECT p.oid,
           p.prokind,                       -- 'f' function, 'p' procedure, 'a' aggregate
           pg_get_function_identity_arguments(p.oid) AS args
    FROM pg_proc p
    JOIN pg_namespace n ON n.oid = p.pronamespace
    WHERE n.nspname = 'public'
      AND NOT EXISTS (
        SELECT 1 FROM pg_depend d
        WHERE d.objid = p.oid
          AND d.deptype = 'e'
      )
  LOOP
    IF r.prokind = 'p' THEN
      EXECUTE format('DROP PROCEDURE IF EXISTS %s(%s) CASCADE', r.oid::regproc, r.args);
    ELSIF r.prokind = 'a' THEN
      EXECUTE format('DROP AGGREGATE IF EXISTS %s(%s) CASCADE', r.oid::regproc, r.args);
    ELSE
      EXECUTE format('DROP FUNCTION IF EXISTS %s(%s) CASCADE', r.oid::regproc, r.args);
    END IF;
  END LOOP;
END$$;`
	if err := RunPSQL(ctx, host, port, user, password, db, dropRoutines); err != nil {
		return err
	}

	// 7) Drop user-defined domains and types (enums/composites) in public
	// Skip extension-owned types and row types of extension-owned views/tables
	dropTypesAndDomains := `
DO $$
DECLARE r record;
BEGIN
  -- Domains (skip extension-owned)
  FOR r IN
    SELECT n.nspname, t.typname, t.oid
    FROM pg_type t
    JOIN pg_namespace n ON n.oid = t.typnamespace
    WHERE n.nspname = 'public' AND t.typtype = 'd'
      AND NOT EXISTS (
        SELECT 1 FROM pg_depend d
        WHERE d.objid = t.oid
          AND d.deptype = 'e'
      )
  LOOP
    EXECUTE format('DROP DOMAIN IF EXISTS %I.%I CASCADE', r.nspname, r.typname);
  END LOOP;

  -- Enums & composite types (exclude array pseudo types, extension-owned, and row types of extension views)
  FOR r IN
    SELECT n.nspname, t.typname, t.oid
    FROM pg_type t
    JOIN pg_namespace n ON n.oid = t.typnamespace
    WHERE n.nspname = 'public'
      AND t.typtype IN ('e','c')
      AND t.typelem = 0
      -- Skip types directly owned by extensions
      AND NOT EXISTS (
        SELECT 1 FROM pg_depend d
        WHERE d.objid = t.oid
          AND d.deptype = 'e'
      )
      -- Skip row types of extension-owned relations (views/tables)
      AND NOT (
        t.typrelid != 0
        AND EXISTS (
          SELECT 1 FROM pg_depend d
          WHERE d.objid = t.typrelid
            AND d.deptype = 'e'
        )
      )
  LOOP
    EXECUTE format('DROP TYPE IF EXISTS %I.%I CASCADE', r.nspname, r.typname);
  END LOOP;
END$$;`
	if err := RunPSQL(ctx, host, port, user, password, db, dropTypesAndDomains); err != nil {
		return err
	}

	return nil
}

func QuoteLiteral(s string) string {
	return `'` + strings.ReplaceAll(s, `'`, `''`) + `'`
}
