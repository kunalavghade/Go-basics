package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	userRepo UserRepo
	mux      *http.ServeMux
}

func main() {
	mux := http.NewServeMux()

	db, err := connectToDB("./users.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	app := &application{
		errorLog: log.New(os.Stderr, "ERROR: ", log.Ltime|log.LstdFlags|log.Lmicroseconds|log.Lshortfile),
		infoLog:  log.New(os.Stdout, "INFO: ", log.Ltime|log.LstdFlags|log.Lmicroseconds|log.Lshortfile),
		userRepo: NewSQLUserRepo(db),
	}
	app.mount(mux)

	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}

func connectToDB(name string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", name)
	return db, err
}
