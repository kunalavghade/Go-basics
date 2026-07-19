package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"kunalavghade/learning-go/section-10/6-repository/repository"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	dbName := "../data.db"
	db, err := connectToDB(dbName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	repo := repository.NewSQLUserRepo(db)
	repo.CreateUser("Kunal Avghade", "kunal@gmail.com", "password123", "https://example.com/avatar.jpg")
	user, _ := repo.GetUserByEmail("kunal@gmail.com")
	userdata, _ := json.MarshalIndent(user, "", "  ")
	fmt.Println(string(userdata))

	PrintUsers(repo)

}

func PrintUsers(repo repository.UserRepo) {
	users, _ := repo.GetUsers()
	userdata, _ := json.MarshalIndent(users, "", "  ")
	fmt.Println("All users: \n", string(userdata))
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

func connectToDB(name string) (*sql.DB, error) {
	db, error := sql.Open("sqlite3", name)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("Connected to DB")

	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
