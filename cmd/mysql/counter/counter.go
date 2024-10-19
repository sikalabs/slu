package counter

import (
	mysqlcmd "github.com/sikalabs/slu/cmd/mysql"
	"github.com/sikalabs/slu/utils/mysql_counter"

	"github.com/spf13/cobra"
)

var CmdFlagName string

var Cmd = &cobra.Command{
	Use:   "counter",
	Short: "Create table _counter and insert one row per second",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		mysql_counter.Counter(
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
