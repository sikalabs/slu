package create

import (
	postgrescmd "github.com/sikalabs/slu/cmd/postgres"
	"github.com/sikalabs/slu/utils/postgres"

	"github.com/spf13/cobra"
)

var PostgresCreateCmdFlagName string

var PostgresCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create Postgres database",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		postgres.Create(
			postgrescmd.PostgresCmdFlagHost,
			postgrescmd.PostgresCmdFlagPort,
			postgrescmd.PostgresCmdFlagUser,
			postgrescmd.PostgresCmdFlagPassword,
			PostgresCreateCmdFlagName,
		)
	},
}

func init() {
	postgrescmd.PostgresCmd.AddCommand(PostgresCreateCmd)
	PostgresCreateCmd.Flags().StringVarP(
		&PostgresCreateCmdFlagName,
		"name",
		"n",
		"",
		"Name of database",
	)
	PostgresCreateCmd.MarkFlagRequired("name")
}
