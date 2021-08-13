package drop

import (
	postgrescmd "github.com/sikalabs/slu/cmd/postgres"
	rootcmd "github.com/sikalabs/slu/cmd/root"
	"github.com/sikalabs/slu/utils/postgres"

	"github.com/spf13/cobra"
)

var PostgresListCmd = &cobra.Command{
	Use:   "list",
	Short: "list Postgres databases",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		if rootcmd.RootCmdFlagJson {
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
}
