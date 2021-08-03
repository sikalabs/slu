package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Create(
	host string,
	port int,
	user string,
	password string,
	name string,
) {
	conninfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s sslmode=disable",
		host, port, user, password,
	)
	db, err := sql.Open("postgres", conninfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Exec("CREATE DATABASE " + name)
	if err != nil {
		fmt.Print(err)

		panic(err)
	}
}
