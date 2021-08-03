package drop

import (
	mysqlcmd "github.com/sikalabs/slut/cmd/mysql"
	"github.com/sikalabs/slut/utils/mysql"

	"github.com/spf13/cobra"
)

var MysqlDropCmdFlagName string

var MysqlDropCmd = &cobra.Command{
	Use:   "drop",
	Short: "Drop MySQL database",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		mysql.Drop(
			mysqlcmd.MysqlCmdFlagHost,
			mysqlcmd.MysqlCmdFlagPort,
			mysqlcmd.MysqlCmdFlagUser,
			mysqlcmd.MysqlCmdFlagPassword,
			MysqlDropCmdFlagName,
		)
	},
}

func init() {
	mysqlcmd.MysqlCmd.AddCommand(MysqlDropCmd)
	MysqlDropCmd.Flags().StringVarP(
		&MysqlDropCmdFlagName,
		"name",
		"n",
		"",
		"Name of database",
	)
	MysqlDropCmd.MarkFlagRequired("name")
}
