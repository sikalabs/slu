package postgres

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/lib/pq"
)

func List(
	host string,
	port int,
	user string,
	password string,
) ([]string, error) {
	db, err := sql.Open(
		"postgres", fmt.Sprintf(
			"host=%s port=%d user=%s password=%s sslmode=disable",
			host, port, user, password,
		))
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT datname FROM pg_database;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dbNames []string
	var dbName string

	for rows.Next() {
		err := rows.Scan(&dbName)
		if err != nil {
			return nil, err
		}
		dbNames = append(dbNames, dbName)
	}
	return dbNames, err
}

func ListText(
	host string,
	port int,
	user string,
	password string,
) {
	dbNames, err := List(host, port, user, password)
	if err != nil {
		panic(err)
	}
	for _, dbName := range dbNames {
		fmt.Println(dbName)
	}
}

func ListJSON(
	host string,
	port int,
	user string,
	password string,
) {
	dbNames, err := List(host, port, user, password)
	if err != nil {
		panic(err)
	}
	outJson, err := json.Marshal(dbNames)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(outJson))
}
