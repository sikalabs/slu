package counter

import (
	postgrescmd "github.com/sikalabs/slu/cmd/postgres"
	"github.com/sikalabs/slu/utils/postgres_counter"

	"github.com/spf13/cobra"
)

var CmdFlagName string

var Cmd = &cobra.Command{
	Use:   "counter",
	Short: "Create table _counter and insert one row per second",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		postgres_counter.Counter(
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
