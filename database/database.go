package database

import (
	"database/sql"
	"fmt"
)

var database *sql.DB

func Initialize() {
	db, connErr := sql.Open("mysql", "root:local@tcp(127.0.0.1:3306)/hedi")

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
