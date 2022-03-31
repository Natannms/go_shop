package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func connectDatabase() *sql.DB {
	conexao := "postgres://postgres:root@localhost:15432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}

	return db
}
