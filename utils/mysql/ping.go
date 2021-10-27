package mysql

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func Ping(
	host string,
	port int,
	user string,
	password string,
	name string,
) {
	db, err := sql.Open(
		"mysql",
		user+":"+password+"@tcp("+host+":"+strconv.Itoa(port)+")/"+name)
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
