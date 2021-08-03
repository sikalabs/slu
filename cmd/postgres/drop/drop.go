package drop

import (
	postgrescmd "github.com/sikalabs/slut/cmd/postgres"
	"github.com/sikalabs/slut/utils/postgres"

	"github.com/spf13/cobra"
)

var PostgresDropCmdFlagName string

var PostgresDropCmd = &cobra.Command{
	Use:   "drop",
	Short: "Drop Postgres database",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		postgres.Drop(
			postgrescmd.PostgresCmdFlagHost,
			postgrescmd.PostgresCmdFlagPort,
			postgrescmd.PostgresCmdFlagUser,
			postgrescmd.PostgresCmdFlagPassword,
			PostgresDropCmdFlagName,
		)
	},
}

func init() {
	postgrescmd.PostgresCmd.AddCommand(PostgresDropCmd)
	PostgresDropCmd.Flags().StringVarP(
		&PostgresDropCmdFlagName,
		"name",
		"n",
		"",
		"Name of database",
	)
	PostgresDropCmd.MarkFlagRequired("name")
}
