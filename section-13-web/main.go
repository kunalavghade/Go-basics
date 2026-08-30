package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golangcollege/sessions"
	_ "github.com/mattn/go-sqlite3"
)

type application struct {
	errorLog    *log.Logger
	infoLog     *log.Logger
	userRepo    UserRepo
	PostRepo    PostRepository
	templateDir string
	publicDir   string
	tp          *TemplateRenderer
	session     *sessions.Session
}

func main() {
	db, err := connectToDB("./data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	session := sessions.New([]byte("ghajchakjshdashdjka"))
	session.Lifetime = 24 * time.Hour
	session.SameSite = http.SameSiteLaxMode

	app := &application{
		errorLog:    log.New(os.Stderr, "ERROR: ", log.Ltime|log.LstdFlags|log.Lmicroseconds|log.Lshortfile),
		infoLog:     log.New(os.Stdout, "INFO: ", log.Ltime|log.LstdFlags|log.Lmicroseconds|log.Lshortfile),
		userRepo:    NewSQLUserRepo(db),
		PostRepo:    NewSQLPostRepository(db),
		templateDir: "./section-13-web/templates",
		publicDir:   "./section-13-web/public",
		session:     session,
	}
	app.tp = NewTemplateRenderer(true, app.templateDir)
	app.infoLog.Println("server running on :8080")
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}

	// comment := "My comment"
	// comments, err := app.PostRepo.GetComments(1)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("%+v\n", comments)

}

func connectToDB(name string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", name)
	return db, err
}
