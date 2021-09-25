package read

import (
	"fmt"

	sqlite_cmd "github.com/sikalabs/slu/cmd/sqlite"
	"github.com/sikalabs/slu/utils/sqlite_utils"
	"github.com/spf13/cobra"
)

var FlagFile string

var Cmd = &cobra.Command{
	Use:   "read",
	Short: "Read SQLite file",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		db, err := sqlite_utils.OpenSQLite(FlagFile)
		if err != nil {
			fmt.Println(err)
		}

		tables := sqlite_utils.GetTables(db)
		for _, table := range tables {
			fmt.Println("Table: ", table)
			sqlite_utils.PrintTable(db, table)
			fmt.Println("")
		}
	},
}

func init() {
	sqlite_cmd.Cmd.AddCommand(Cmd)
	Cmd.Flags().StringVarP(
		&FlagFile,
		"file",
		"f",
		"",
		"SQLite file",
	)
	Cmd.MarkFlagRequired("file")
}
