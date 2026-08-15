package main

import (
	"net/http"
)

const htmlContent = `
<!DOCTYPE html>
<html>
<head>
	<title>%s</title>
</head>
<body>
	%s
</body>
</html>
`

func (app *application) errorpage(w http.ResponseWriter, r *http.Request) {
	panic("Helo")
}

func (app *application) submit(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Printf("user Id : %v", app.session.Get(r, "userId"))
	app.render(w, r, "submit.html", nil)
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Printf("user Id : %v", app.session.Get(r, "userId"))
	app.render(w, r, "index.html", nil)
}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "about.html", nil)
}

func (app *application) contact(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "contact.html", nil)
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	// app.session.Put(r, "userId", 123)
	// data := map[string]interface{}{
	// 	"Error":    "",
	// 	"Username": "",
	// 	"Password": "",
	// }
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		form := NewForm(r.PostForm)
		form.required("email", "password").maxLength("email", 50).minLength("password", 6)
		if !form.valid() {
			app.errorLog.Printf("Invalid form data %v", form.Errors)
			app.render(w, r, "login.html", nil)
			return
		}

		email := r.FormValue("email")
		password := r.FormValue("password")

		app.infoLog.Printf("Logged in with email %s and password %s", email, password)
	}
	app.render(w, r, "login.html", nil)
}

func (app *application) register(w http.ResponseWriter, r *http.Request) {
	// data := map[string]interface{}{
	// 	"Error":           "",
	// 	"Username":        "",
	// 	"Email":           "",
	// 	"Password":        "",
	// 	"ConfirmPassword": "",
	// }
	// app.render(w, r, "register.html", data)
}
