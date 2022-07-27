package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var database *sql.DB

func Initialize() {
	db, connErr := sql.Open("sqlite3", "./database/database.db")

	db.SetMaxOpenConns(8)
	db.SetMaxIdleConns(8)

	if connErr != nil {
		fmt.Printf("MySQL connetion could not be established...\n")

		return
	}

	database = db
}

func Deinitialize() {
	database.Close()
}
