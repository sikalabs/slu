package drop

import (
	mysqlcmd "github.com/sikalabs/slu/cmd/mysql"
	"github.com/sikalabs/slu/utils/mysql"

	"github.com/spf13/cobra"
)

var FlagJson bool

var MysqlListCmd = &cobra.Command{
	Use:   "list",
	Short: "list Mysql databases",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		if FlagJson {
			mysql.ListJSON(
				mysqlcmd.MysqlCmdFlagHost,
				mysqlcmd.MysqlCmdFlagPort,
				mysqlcmd.MysqlCmdFlagUser,
				mysqlcmd.MysqlCmdFlagPassword,
			)
		} else {
			mysql.ListText(
				mysqlcmd.MysqlCmdFlagHost,
				mysqlcmd.MysqlCmdFlagPort,
				mysqlcmd.MysqlCmdFlagUser,
				mysqlcmd.MysqlCmdFlagPassword,
			)
		}
	},
}

func init() {
	mysqlcmd.MysqlCmd.AddCommand(MysqlListCmd)
	MysqlListCmd.PersistentFlags().BoolVar(
		&FlagJson,
		"json",
		false,
		"Format output to JSON",
	)
}
