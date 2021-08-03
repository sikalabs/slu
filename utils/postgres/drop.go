package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Drop(
	host string,
	port int,
	user string,
	password string,
	name string,
) {
	db, err := sql.Open(
		"postgres", fmt.Sprintf(
			"host=%s port=%d user=%s password=%s sslmode=disable",
			host, port, user, password,
		))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("DROP DATABASE " + name)
	if err != nil {
		panic(err)
	}
}
