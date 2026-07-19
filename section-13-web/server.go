package main

import (
	"net/http"
)

func (app *application) Serve() error {
	srv := http.Server{
		Addr:    ":8080",
		Handler: app.routes(),
	}
	return srv.ListenAndServe()
}
