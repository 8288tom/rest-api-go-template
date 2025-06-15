package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// basic sqlite implementation - remember to download driver when changing DB (this one is a simple file based DB)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {

	createTemplateTable := `
	CREATE TABLE IF NOT EXISTS template(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
		
	)`

	_, err := DB.Exec(createTemplateTable)

	if err != nil {
		fmt.Print(err)
		panic("Could not create users table")
	}
}
