package mysql

import (
	"database/sql"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func Create(
	host string,
	port int,
	user string,
	password string,
	name string,
) {
	db, err := sql.Open(
		"mysql",
		user+":"+password+"@tcp("+host+":"+strconv.Itoa(port)+")/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE " + name)
	if err != nil {
		panic(err)
	}
}
