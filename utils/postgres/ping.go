package postgres

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func Ping(
	host string,
	port int,
	user string,
	password string,
	name string,
) {
	conninfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, name,
	)
	db, err := sql.Open("postgres", conninfo)
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
