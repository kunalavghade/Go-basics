package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
)

type contextKey string

const contextAuthKey contextKey = "isAuthenticated"

func (app *application) logger(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.infoLog.Printf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func (app *application) recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				app.serverError(w, fmt.Errorf("%s", err))
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func (app *application) requireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !app.isAthenticated(r) {
			app.infoLog.Println("Unauthenticated access attempt")
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (app *application) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		exist := app.session.Exists(r, loggedInUserKey)
		if !exist {
			next.ServeHTTP(w, r)
			return
		}
		_, err := app.userRepo.GetUserByEmail(app.session.GetString(r, loggedInUserKey))
		if errors.Is(err, sql.ErrNoRows) {
			app.session.Remove(r, loggedInUserKey)
			next.ServeHTTP(w, r)
			return
		} else if err != nil {
			app.serverError(w, err)
			return
		}

		ctx := context.WithValue(r.Context(), contextAuthKey, true)
		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}

func (app *application) isAthenticated(r *http.Request) bool {
	return app.session.Exists(r, loggedInUserKey)
}
