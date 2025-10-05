package postgres

import (
	"compress/gzip"
	"context"
	"fmt"
	"io"

	"github.com/sikalabs/slu/utils/azure_utils"
	"github.com/sikalabs/slu/utils/postgres_utils"
)

func RestoreFromAzureBlob(
	host string,
	port int,
	user string,
	password string,
	dbName string,
	containerName string,
	accountName string,
	accountKey string,
	sourcePrefix string,
	sourceSuffix string,
) error {
	ctx := context.Background()

	fmt.Println("Creating Azure Blob client…")
	client, err := azure_utils.AzureBlobClient(accountName, accountKey)
	if err != nil {
		return fmt.Errorf("azure client: %w", err)
	}

	fmt.Println("Selecting latest blob by prefix/suffix…")
	blobName, blobSize, err := azure_utils.FindLatestBlob(ctx, client, containerName, sourcePrefix, sourceSuffix)
	if err != nil {
		return fmt.Errorf("find latest blob: %w", err)
	}
	fmt.Printf("Selected blob: %s (%.2f MB)\n", blobName, float64(blobSize)/1024/1024)

	fmt.Println("Opening streaming download from Azure…")
	body, closer, err := azure_utils.OpenBlobStream(ctx, client, containerName, blobName)
	if err != nil {
		return fmt.Errorf("open blob stream: %w", err)
	}
	defer closer()

	fmt.Println("Setting up on-the-fly gunzip…")
	sqlReader, gzClose, err := streamGunzip(body)
	if err != nil {
		return fmt.Errorf("gunzip stream: %w", err)
	}
	defer gzClose()

	// Use admin DB "postgres" to manage grants & terminate sessions
	adminDB := "postgres"

	fmt.Println("Restricting connections (allow only provided user)…")
	if err := postgres_utils.RestrictPSQLConnections(ctx, host, port, user, password, adminDB, dbName, user); err != nil {
		return fmt.Errorf("restrict connections: %w", err)
	}
	fmt.Println("Connections restricted and foreign sessions terminated.")

	fmt.Println("Purging ALL objects in schema public (views, tables, sequences, routines, types)…")
	if err := postgres_utils.PurgePSQLPublicSchemaObjects(ctx, host, port, user, password, dbName); err != nil {
		return fmt.Errorf("purge public schema: %w", err)
	}
	fmt.Println("public schema emptied (schema itself preserved).")

	fmt.Println("Restoring database via psql (streaming gzip→sql→stdin)…")
	if err := postgres_utils.RunPSQLFromReader(ctx, host, port, user, password, dbName, sqlReader); err != nil {
		return fmt.Errorf("psql restore stream: %w", err)
	}
	fmt.Println("Restore completed successfully.")

	fmt.Println("Re-enabling connections for PUBLIC…")
	if err := postgres_utils.ReopenPSQLConnections(ctx, host, port, user, password, adminDB, dbName); err != nil {
		return fmt.Errorf("reopen connections: %w", err)
	}
	fmt.Println("Connections re-enabled.")

	fmt.Println("Done.")
	return nil
}

func streamGunzip(r io.Reader) (io.Reader, func(), error) {
	gzr, err := gzip.NewReader(r)
	if err != nil {
		return nil, nil, err
	}
	closer := func() { _ = gzr.Close() }
	return gzr, closer, nil
}
