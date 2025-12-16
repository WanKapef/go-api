package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectSQLite(path string) *sql.DB {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		log.Fatal(err)
	}

	// SQLite precisa disso para evitar erro "database is locked"
	db.SetMaxOpenConns(1)

	return db
}
