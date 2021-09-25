package read

import (
	"fmt"
	"log"
	"os"

	sqlite_cmd "github.com/sikalabs/slu/cmd/sqlite"
	"github.com/sikalabs/slu/utils/sqlite_utils"
	"github.com/spf13/cobra"
)

var FlagFile string
var FlagTable string

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

		if FlagTable == "" {
			for _, table := range tables {
				fmt.Println("Table: ", table)
				sqlite_utils.PrintTable(db, table)
				fmt.Println("")
			}
		} else {
			for _, table := range tables {
				if table == FlagTable {
					fmt.Println("Table: ", table)
					sqlite_utils.PrintTable(db, table)
					os.Exit(0)
				}
			}
			log.Fatal(fmt.Errorf("table %s not found", FlagTable))
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
	Cmd.Flags().StringVarP(
		&FlagTable,
		"table",
		"t",
		"",
		"SQLite table",
	)
}
