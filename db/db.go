package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	// DB, err := sql.Open("sqlite3", "api.db")
	var err error
	DB, err = sql.Open("sqlite3", "./db/api.db")
	if err != nil {
		// panic("Could not connect to the DB...")
		fmt.Println("error from the initDB", err.Error())
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTables()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	
	)`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		fmt.Println(err.Error())
		panic("could not create user table")
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY(user_id) REFERENCES users(id)
	)`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		// panic("Could not create events table...")
		fmt.Println("error from the createTable()", err.Error())
	}
}
