package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	defaultMiddleware := alice.New(app.logger, app.recover)
	secureMiddleware := alice.New(app.session.Enable, app.authenticate)

	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir(app.publicDir))))

	mux.Handle("/", secureMiddleware.Then(http.HandlerFunc(app.home)))
	mux.Handle("/login", secureMiddleware.Then(http.HandlerFunc(app.login)))
	mux.Handle("/logout", secureMiddleware.Then(http.HandlerFunc(app.logout)))
	mux.Handle("/submit", secureMiddleware.Append(app.requireAuth).Then(http.HandlerFunc(app.submit)))
	mux.Handle("/register", secureMiddleware.Then(http.HandlerFunc(app.register)))
	mux.Handle("/error", secureMiddleware.Then(http.HandlerFunc(app.errorpage)))
	mux.HandleFunc("/about", app.about)
	mux.HandleFunc("/contact", app.contact)

	// handler := app.recover(app.logger(app.session.Enable((mux))))
	// return handler
	return defaultMiddleware.Then(mux)
}
