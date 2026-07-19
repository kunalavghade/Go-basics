package main

import (
	"fmt"
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
	pageTitle := "Home"
	pageContent := "Welcome to my website"
	fmt.Fprintf(w, htmlContent, pageTitle, pageContent)
}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	pageTitle := "About"
	pageContent := "About page"
	fmt.Fprintf(w, htmlContent, pageTitle, pageContent)
}

func (app *application) contact(w http.ResponseWriter, r *http.Request) {
	pageTitle := "Contact"
	pageContent := "Contact page"
	fmt.Fprintf(w, htmlContent, pageTitle, pageContent)
}
