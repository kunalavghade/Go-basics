package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	dbName := "../data.db"
	// _ = os.Remove(dbName)

	db, error := sql.Open("sqlite3", dbName)
	if error != nil {
		log.Fatal(error)
	}
	defer db.Close()
	fmt.Println("Connected to DB")

	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	creatUser(db, "Kunal Avghade", "[EMAIL_ADDRESS]", "password")

}

func creatUser(db *sql.DB, name, email, password string) {
	var query = `INSERT INTO user(name, email, hashed_password) VALUES(?, ?, ?)`
	_, err := db.Exec(query, name, email, password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("User created successfully")
}
