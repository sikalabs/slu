package drop

import (
	postgrescmd "github.com/sikalabs/slu/cmd/postgres"
	"github.com/sikalabs/slu/utils/postgres"

	"github.com/spf13/cobra"
)

var FlagJson bool

var PostgresListCmd = &cobra.Command{
	Use:   "list",
	Short: "list Postgres databases",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		if FlagJson {
			postgres.ListJSON(
				postgrescmd.PostgresCmdFlagHost,
				postgrescmd.PostgresCmdFlagPort,
				postgrescmd.PostgresCmdFlagUser,
				postgrescmd.PostgresCmdFlagPassword,
			)
		} else {
			postgres.ListText(
				postgrescmd.PostgresCmdFlagHost,
				postgrescmd.PostgresCmdFlagPort,
				postgrescmd.PostgresCmdFlagUser,
				postgrescmd.PostgresCmdFlagPassword,
			)
		}
	},
}

func init() {
	postgrescmd.PostgresCmd.AddCommand(PostgresListCmd)
	PostgresListCmd.PersistentFlags().BoolVar(
		&FlagJson,
		"json",
		false,
		"Format output to JSON",
	)
}
