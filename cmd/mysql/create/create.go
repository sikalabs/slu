package create

import (
	mysqlcmd "github.com/sikalabs/slut/cmd/mysql"
	"github.com/sikalabs/slut/utils/mysql"

	"github.com/spf13/cobra"
)

var MysqlCreateCmdFlagName string

var MysqlCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create MtSQL database",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		mysql.Create(
			mysqlcmd.MysqlCmdFlagHost,
			mysqlcmd.MysqlCmdFlagPort,
			mysqlcmd.MysqlCmdFlagUser,
			mysqlcmd.MysqlCmdFlagPassword,
			MysqlCreateCmdFlagName,
		)
	},
}

func init() {
	mysqlcmd.MysqlCmd.AddCommand(MysqlCreateCmd)
	MysqlCreateCmd.Flags().StringVarP(
		&MysqlCreateCmdFlagName,
		"name",
		"n",
		"",
		"Name of database",
	)
	MysqlCreateCmd.MarkFlagRequired("name")
}
