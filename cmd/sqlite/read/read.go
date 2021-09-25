package read

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	sqlite_cmd "github.com/sikalabs/slu/cmd/sqlite"
	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var FlagFile string

var Cmd = &cobra.Command{
	Use:   "read",
	Short: "Read SQLite file",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		db, err := gorm.Open(sqlite.Open(FlagFile), &gorm.Config{})
		if err != nil {
			fmt.Println(err)
		}

		var tables []map[string]interface{}
		db.Raw("SELECT distinct tbl_name from sqlite_master order by 1").Scan(&tables)
		for _, el := range tables {
			table := el["tbl_name"].(string)
			fmt.Println("Table: ", table)
			tw := tablewriter.NewWriter(os.Stdout)

			var columns []string
			db.Raw("SELECT name FROM PRAGMA_TABLE_INFO('" + table + "')").Scan(&columns)

			tw.SetHeader(columns)

			var data []map[string]interface{}
			db.Raw("SELECT * from " + table + " order by 1").Scan(&data)

			for _, row := range data {
				var y []string
				for _, colName := range columns {
					colValue := row[colName]
					var s string
					s = fmt.Sprintf("%v", colValue)
					if colName == "created_at" && s != "<nil>" {
						s = s[0:19]
					} else if colName == "updated_at" && s != "<nil>" {
						s = s[0:19]
					} else if colName == "deleted_at" && s != "<nil>" {
						s = s[0:19]
					}
					y = append(y, s)
				}
				tw.Append(y)
			}
			tw.Render()
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
