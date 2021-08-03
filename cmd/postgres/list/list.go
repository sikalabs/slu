package drop

import (
	postgrescmd "github.com/sikalabs/slut/cmd/postgres"
	"github.com/sikalabs/slut/utils/postgres"

	"github.com/spf13/cobra"
)

var PostgresListCmd = &cobra.Command{
	Use:   "list",
	Short: "list Postgres databases",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		postgres.ListText(
			postgrescmd.PostgresCmdFlagHost,
			postgrescmd.PostgresCmdFlagPort,
			postgrescmd.PostgresCmdFlagUser,
			postgrescmd.PostgresCmdFlagPassword,
		)
	},
}

func init() {
	postgrescmd.PostgresCmd.AddCommand(PostgresListCmd)
}
