package main

import (
	"errors"
	"net/http"
)

func (app *application) Serve() error {
	if app.mux == nil {
		return errors.New("Mux is not initialized")
	}
	return http.ListenAndServe(":8080", app.mux)
}

func (app *application) mount(mux *http.ServeMux) {
	app.mux = mux
	app.mux.HandleFunc("/", app.home)
	app.mux.HandleFunc("/about", app.about)
	app.mux.HandleFunc("/contact", app.contact)
}
