package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var table = `
CREATE TABLE IF NOT EXISTS user (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	email TEXT NOT NULL,
	hashed_password BLOB NOT NULL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP
)
`

func main() {
	dbName := "../data.db"
	_ = os.Remove(dbName)

	db, error := sql.Open("sqlite3", dbName)
	if error != nil {
		log.Fatal(error)
	}
	defer db.Close()
	fmt.Println("Connected to DB")

	// absPath, err := filepath.Abs(dbName)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Location of db full path is %v\n", absPath)

	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(table)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB is ready")
}
