package db

import (
	"database/sql"

	_ "github.com/glebarez/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "api.db")

	if err != nil {
		panic("Could not connect to the database." + err.Error())
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
		panic("Could not create users table: " + err.Error())
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
		panic("Could not create events table: " + err.Error())
	}

}
