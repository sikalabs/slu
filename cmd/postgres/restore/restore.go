package restore

import (
	postgrescmd "github.com/sikalabs/slu/cmd/postgres"
	"github.com/spf13/cobra"
)

var PostgresRestoreFlagName string

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
}
