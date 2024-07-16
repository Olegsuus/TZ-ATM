package db

import (
	"log"

	"database/sql"
)

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", "file:accounts.db?cache=shared&mode=rwc")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS accounts (id TEXT PRIMARY KEY, balance REAL)")
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}
	return db
}
