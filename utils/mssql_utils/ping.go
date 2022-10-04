package mssql_utils

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/denisenkom/go-mssqldb/azuread"
)

func Ping(
	host string,
	port int,
	user string,
	password string,
) {
	db, err := sql.Open(
		azuread.DriverName,
		"sqlserver://"+user+":"+password+"@"+host+":"+strconv.Itoa(port))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("OK")
	os.Exit(0)
}
