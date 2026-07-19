package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type user struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

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

	id, _ := cratePreparedUser(db, "kunal", "[EMAIL_ADDRESS]", "dfghjkl;")
	fmt.Println(id)

	// user, err := getUserByID(db, 1)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(user)

	// users, _ := getUsers(db)
	// fmt.Println(users)

}

func cratePreparedUser(db *sql.DB, name, email, password string) (int, error) {
	query := `insert into user(name, email, hashed_password) values (?, ?, ?)`
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(name, email, password)
	if err != nil {
		log.Fatal(err)
	}
	id, _ := result.LastInsertId()
	return int(id), nil
}

func getUserByID(db *sql.DB, id int) (user, error) {
	query := `SELECT id, name, email, hashed_password, created_at FROM user WHERE id = ?`
	row := db.QueryRow(query, id)

	var usr user

	err := row.Scan(
		&usr.ID,
		&usr.Name,
		&usr.Email,
		&usr.Password,
		&usr.CreatedAt,
	)

	if err != nil {
		log.Fatal(err)
	}
	return usr, nil
}

func getUsers(db *sql.DB) ([]user, error) {
	query := `SELECT id, name, email, hashed_password, created_at FROM user`
	row, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	var users []user
	for row.Next() {
		var usr user
		err := row.Scan(
			&usr.ID,
			&usr.Name,
			&usr.Email,
			&usr.Password,
			&usr.CreatedAt,
		)

		if err != nil {
			log.Fatal(err)
		}
		users = append(users, usr)
	}
	return users, nil
}
