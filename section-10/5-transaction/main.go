package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var table = `
CREATE TABLE IF NOT EXISTS profile (
	user_id INTEGER PRIMARY KEY REFERENCES user(id) ON DELETE CASCADE,
	avatar TEXT NOT NULL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP
)
`

type Profile struct {
	UserID int    `json:"user_id"`
	Avatar string `json:"avatar"`
}

type user struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	Profile   Profile   `json:"profile"`
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

	createTable(db)
	id, err := cratePreparedUser(db, "josh", "test@tes1t.com", "dfghjkl", "http://avatar.com/kunal")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(id)

	user, err := getUserByEmail(db, "test@tes1t.com")
	if err != nil {
		log.Fatal(err)
	}
	data, err := json.MarshalIndent(user, " ", " ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))

	// users, _ := getUsers(db)
	// fmt.Println(users)

}

func createTable(db *sql.DB) error {
	_, err := db.Exec(table)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func cratePreparedUser(db *sql.DB, name, email, password, avatar string) (int, error) {
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	query := `insert into user(name, email, hashed_password) values (?, ?, ?)`
	stmt, err := tx.PrepareContext(ctx, query)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(name, email, password)
	if err != nil {
		log.Fatal(err)
		return 0, err
	}
	id, _ := result.LastInsertId()

	profileStmt, err := tx.PrepareContext(ctx, `insert into profile(user_id, avatar) values (?,?)`)
	if err != nil {
		log.Fatal(err)
		err = tx.Rollback()
		return 0, err
	}
	defer profileStmt.Close()

	_, err = profileStmt.Exec(id, avatar)
	if err != nil {
		err = tx.Rollback()
		log.Fatal(err)
		return 0, err
	}
	return int(id), tx.Commit()
}

func getUserByEmail(db *sql.DB, email string) (user, error) {
	query := `
		SELECT 
			u.id, u.name, u.email, u.hashed_password, u.created_at, p.user_id, p.avatar
		FROM user u
		JOIN profile p
		ON u.id = p.user_id
		WHERE u.email = ?`
	row := db.QueryRow(query, email)

	var usr user

	err := row.Scan(
		&usr.ID,
		&usr.Name,
		&usr.Email,
		&usr.Password,
		&usr.CreatedAt,
		&usr.Profile.UserID,
		&usr.Profile.Avatar,
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
