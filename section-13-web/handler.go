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

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.render(w, "index.html", nil)
}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	app.render(w, "about.html", nil)
}

func (app *application) contact(w http.ResponseWriter, r *http.Request) {
	app.render(w, "contact.html", nil)
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Error":    "",
		"Username": "",
		"Password": "",
	}
	app.render(w, "login.html", data)
}

func (app *application) register(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Error":           "",
		"Username":        "",
		"Email":           "",
		"Password":        "",
		"ConfirmPassword": "",
	}
	app.render(w, "register.html", data)
}
