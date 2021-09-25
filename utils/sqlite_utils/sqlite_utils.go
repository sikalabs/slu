package sqlite_utils

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func OpenSQLite(file string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(file), &gorm.Config{})
	return db, err
}

func GetTables(db *gorm.DB) []string {
	var tables []string
	db.
		Raw("SELECT distinct tbl_name from sqlite_master order by 1").
		Scan(&tables)
	return tables
}

func GetColumns(db *gorm.DB, table string) []string {
	var columns []string
	db.
		Raw("SELECT name FROM PRAGMA_TABLE_INFO('" + table + "')").
		Scan(&columns)
	return columns
}

func process(column string, data interface{}) string {
	var s string
	s = fmt.Sprintf("%v", data)
	if column == "created_at" && s != "<nil>" {
		s = s[0:19]
	} else if column == "updated_at" && s != "<nil>" {
		s = s[0:19]
	} else if column == "deleted_at" && s != "<nil>" {
		s = s[0:19]
	}
	return s
}

func GetRows(db *gorm.DB, table string) [][]string {
	var data [][]string

	columns := GetColumns(db, table)

	var rawData []map[string]interface{}
	db.Raw("SELECT * from " + table + " order by 1").Scan(&rawData)

	for _, rawRow := range rawData {
		var row []string
		for _, colName := range columns {
			colValue := rawRow[colName]
			row = append(row, process(colName, colValue))
		}
		data = append(data, row)
	}

	return data
}

func PrintTable(db *gorm.DB, table string) {
	tw := tablewriter.NewWriter(os.Stdout)
	tw.SetHeader(GetColumns(db, table))
	data := GetRows(db, table)

	for _, row := range data {
		tw.Append(row)
	}

	tw.Render()
}
