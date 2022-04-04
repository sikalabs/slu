package mysql

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var MysqlCmdFlagUser string
var MysqlCmdFlagPassword string
var MysqlCmdFlagHost string
var MysqlCmdFlagPort int

var MysqlCmd = &cobra.Command{
	Use:   "mysql",
	Short: "MySQL Utils",
}

func init() {
	root.RootCmd.AddCommand(MysqlCmd)
	MysqlCmd.PersistentFlags().StringVarP(
		&MysqlCmdFlagHost,
		"host",
		"H",
		"127.0.0.1",
		"Database host",
	)
	MysqlCmd.PersistentFlags().IntVarP(
		&MysqlCmdFlagPort,
		"port",
		"P",
		3306,
		"Database port",
	)
	MysqlCmd.PersistentFlags().StringVarP(
		&MysqlCmdFlagUser,
		"user",
		"u",
		"",
		"Database user",
	)
	MysqlCmd.MarkPersistentFlagRequired("user")
	MysqlCmd.PersistentFlags().StringVarP(
		&MysqlCmdFlagPassword,
		"password",
		"p",
		"",
		"Database password",
	)
	MysqlCmd.MarkPersistentFlagRequired("password")
}
