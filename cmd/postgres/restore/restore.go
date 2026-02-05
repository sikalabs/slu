package restore

import (
	postgrescmd "github.com/sikalabs/slu/cmd/postgres"
	"github.com/spf13/cobra"
)

var PostgresRestoreFlagName string
var PostgresRestoreFlagSSLMode string

var PostgresRestoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Postgres Restore Utils",
}

func init() {
	postgrescmd.PostgresCmd.AddCommand(PostgresRestoreCmd)
	PostgresRestoreCmd.PersistentFlags().StringVarP(
		&PostgresRestoreFlagName,
		"name",
		"n",
		"",
		"Name of database to restore",
	)
	PostgresRestoreCmd.MarkPersistentFlagRequired("name")
	PostgresRestoreCmd.PersistentFlags().StringVarP(
		&PostgresRestoreFlagSSLMode,
		"ssl-mode",
		"S",
		"disable",
		"SSL mode (disable, allow, prefer, require, verify-ca, verify-full)",
	)
}
