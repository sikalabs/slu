package ping

import (
	mysqlcmd "github.com/sikalabs/slu/cmd/mysql"
	"github.com/sikalabs/slu/utils/mysql"

	"github.com/spf13/cobra"
)

var CmdFlagName string

var Cmd = &cobra.Command{
	Use:   "ping",
	Short: "Test connection to MySQL database",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		mysql.Ping(
			mysqlcmd.MysqlCmdFlagHost,
			mysqlcmd.MysqlCmdFlagPort,
			mysqlcmd.MysqlCmdFlagUser,
			mysqlcmd.MysqlCmdFlagPassword,
			CmdFlagName,
		)
	},
}

func init() {
	mysqlcmd.MysqlCmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&CmdFlagName,
		"name",
		"n",
		"",
		"Name of database",
	)
	Cmd.MarkFlagRequired("name")
}
