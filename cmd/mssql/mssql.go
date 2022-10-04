package mssql

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var MssqlCmdFlagUser string
var MssqlCmdFlagPassword string
var MssqlCmdFlagHost string
var MssqlCmdFlagPort int

var MssqlCmd = &cobra.Command{
	Use:   "mssql",
	Short: "MSSQL Utils (Microsoft SQL Server)",
}

func init() {
	root.RootCmd.AddCommand(MssqlCmd)
	MssqlCmd.PersistentFlags().StringVarP(
		&MssqlCmdFlagHost,
		"host",
		"H",
		"127.0.0.1",
		"Database host",
	)
	MssqlCmd.PersistentFlags().IntVarP(
		&MssqlCmdFlagPort,
		"port",
		"P",
		1433,
		"Database port",
	)
	MssqlCmd.PersistentFlags().StringVarP(
		&MssqlCmdFlagUser,
		"user",
		"u",
		"",
		"Database user",
	)
	MssqlCmd.MarkPersistentFlagRequired("user")
	MssqlCmd.PersistentFlags().StringVarP(
		&MssqlCmdFlagPassword,
		"password",
		"p",
		"",
		"Database password",
	)
	MssqlCmd.MarkPersistentFlagRequired("password")
}
