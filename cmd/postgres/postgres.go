package postgres

import (
	"github.com/sikalabs/slu/cmd/root"
	"github.com/spf13/cobra"
)

var PostgresCmdFlagUser string
var PostgresCmdFlagPassword string
var PostgresCmdFlagHost string
var PostgresCmdFlagPort int

var PostgresCmd = &cobra.Command{
	Use:   "postgres",
	Short: "Postgres Utils",
}

func init() {
	root.RootCmd.AddCommand(PostgresCmd)
	PostgresCmd.PersistentFlags().StringVarP(
		&PostgresCmdFlagHost,
		"host",
		"H",
		"",
		"Database host",
	)
	PostgresCmd.MarkPersistentFlagRequired("host")
	PostgresCmd.PersistentFlags().IntVarP(
		&PostgresCmdFlagPort,
		"port",
		"P",
		0,
		"Database port",
	)
	PostgresCmd.MarkPersistentFlagRequired("port")
	PostgresCmd.PersistentFlags().StringVarP(
		&PostgresCmdFlagUser,
		"user",
		"u",
		"",
		"Database user",
	)
	PostgresCmd.MarkPersistentFlagRequired("user")
	PostgresCmd.PersistentFlags().StringVarP(
		&PostgresCmdFlagPassword,
		"password",
		"p",
		"",
		"Database password",
	)
	PostgresCmd.MarkPersistentFlagRequired("password")
}
