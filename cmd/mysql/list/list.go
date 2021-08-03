package drop

import (
	mysqlcmd "github.com/sikalabs/slut/cmd/mysql"
	rootcmd "github.com/sikalabs/slut/cmd/root"
	"github.com/sikalabs/slut/utils/mysql"

	"github.com/spf13/cobra"
)

var MysqlListCmd = &cobra.Command{
	Use:   "list",
	Short: "list Mysql databases",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		if rootcmd.RootCmdFlagJson {
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
}
