package ping

import (
	postgrescmd "github.com/sikalabs/slu/cmd/postgres"
	"github.com/sikalabs/slu/utils/postgres"

	"github.com/spf13/cobra"
)

var CmdFlagName string

var Cmd = &cobra.Command{
	Use:   "ping",
	Short: "Test connection to Postgres database",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		postgres.Ping(
			postgrescmd.PostgresCmdFlagHost,
			postgrescmd.PostgresCmdFlagPort,
			postgrescmd.PostgresCmdFlagUser,
			postgrescmd.PostgresCmdFlagPassword,
			CmdFlagName,
		)
	},
}

func init() {
	postgrescmd.PostgresCmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&CmdFlagName,
		"name",
		"n",
		"",
		"Name of database",
	)
	Cmd.MarkFlagRequired("name")
}
