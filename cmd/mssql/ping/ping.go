package ping

import (
	mssqlcmd "github.com/sikalabs/slu/cmd/mssql"
	"github.com/sikalabs/slu/utils/mssql_utils"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "ping",
	Short: "Test connection to MySQL database",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		mssql_utils.Ping(
			mssqlcmd.MssqlCmdFlagHost,
			mssqlcmd.MssqlCmdFlagPort,
			mssqlcmd.MssqlCmdFlagUser,
			mssqlcmd.MssqlCmdFlagPassword,
		)
	},
}

func init() {
	mssqlcmd.MssqlCmd.AddCommand(Cmd)
}
