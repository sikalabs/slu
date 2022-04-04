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
		"127.0.0.1",
		"Database host",
	)
	PostgresCmd.PersistentFlags().IntVarP(
		&PostgresCmdFlagPort,
		"port",
		"P",
		5432,
		"Database port",
	)
	PostgresCmd.PersistentFlags().StringVarP(
		&PostgresCmdFlagUser,
		"user",
		"u",
		"postgres",
		"Database user",
	)
	PostgresCmd.PersistentFlags().StringVarP(
		&PostgresCmdFlagPassword,
		"password",
		"p",
		"",
		"Database password",
	)
	PostgresCmd.MarkPersistentFlagRequired("password")
}
