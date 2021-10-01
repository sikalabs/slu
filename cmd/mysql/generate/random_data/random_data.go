package random_data

import (
	mysql_cmd "github.com/sikalabs/slu/cmd/mysql"
	mysql_generate_cmd "github.com/sikalabs/slu/cmd/mysql/generate"
	"github.com/sikalabs/slu/utils/mysql_random_utils"

	"github.com/spf13/cobra"
)

var FlagName string
var FlagBatchSize int
var FlagBatchCount int

var Cmd = &cobra.Command{
	Use:     "random-data",
	Short:   "Drop MySQL database",
	Aliases: []string{"r", "ran", "rnd"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		mysql_random_utils.GenerateRandomData(
			mysql_cmd.MysqlCmdFlagHost,
			mysql_cmd.MysqlCmdFlagPort,
			mysql_cmd.MysqlCmdFlagUser,
			mysql_cmd.MysqlCmdFlagPassword,
			FlagName,
			FlagBatchSize,
			FlagBatchCount,
		)
	},
}

func init() {
	mysql_generate_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagName,
		"name",
		"n",
		"",
		"Name of database",
	)
	Cmd.MarkFlagRequired("name")
	Cmd.Flags().IntVar(
		&FlagBatchSize,
		"batch-size",
		100,
		"Nuber of rows in one SQL insert",
	)
	Cmd.Flags().IntVar(
		&FlagBatchCount,
		"batch-count",
		1000,
		"Count of inserts",
	)
}
