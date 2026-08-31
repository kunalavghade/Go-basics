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

const loggedInUserKey = `logged_user_id`

func (app *application) errorpage(w http.ResponseWriter, r *http.Request) {
	panic("Helo")
}

func (app *application) submit(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Printf("user Id : %v", app.session.Get(r, loggedInUserKey))
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		form := NewForm(r.PostForm)
		form.required("title", "url").maxLength("title", 100).maxLength("url", 500)
		if !form.valid() {
			form.Errors.Add("generic", "Invalid Details Provided..")
			app.render(w, r, "submit.html", &templateData{
				Form: form,
			})
			return
		}
		url := r.FormValue("url")
		title := r.FormValue("title")
		user, err := app.getUserFromCtx(r.Context())
		if err != nil {
			form.Errors.Add("generic", err.Error())
			app.render(w, r, "submit.html", &templateData{
				Form: form,
			})
			return
		}
		id, err := app.PostRepo.CreatePost(title, url, user.ID)
		if err != nil {
			form.Errors.Add("generic", err.Error())
			app.render(w, r, "submit.html", &templateData{
				Form: form,
			})
			return
		}

		app.session.Put(r, "flash", "Post created successfully!")
		app.infoLog.Printf("Post create id %d", id)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	app.render(w, r, "submit.html", &templateData{
		Form: NewForm(r.PostForm),
	})
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Printf("user Id : %v", app.session.Get(r, loggedInUserKey))
	app.render(w, r, "index.html", nil)
}

func (app *application) about(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "about.html", nil)
}

func (app *application) contact(w http.ResponseWriter, r *http.Request) {
	app.render(w, r, "contact.html", nil)
}

func (app *application) login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		form := NewForm(r.PostForm)
		form.required("email", "password").maxLength("email", 50).minLength("password", 6)
		if !form.valid() {
			app.errorLog.Printf("Invalid form data %v", form.Errors)
			form.Errors.Add("generic", "Invalid credentials")
			app.render(w, r, "login.html", &templateData{
				Form: form,
			})
			return
		}

		email := r.FormValue("email")
		password := r.FormValue("password")

		_, err := app.userRepo.AuthenticateUser(email, password)
		if err != nil {
			form.Errors.Add("generic", err.Error())
			app.render(w, r, "login.html", &templateData{
				Form: form,
			})
			return
		}

		// logged in
		app.session.Put(r, loggedInUserKey, email)
		app.session.Put(r, "flash", "Your logged In !")

		app.infoLog.Printf("Logged in successfully")
		http.Redirect(w, r, "/submit", http.StatusSeeOther)
		return
	}
	app.render(w, r, "login.html", &templateData{
		Form: NewForm(r.PostForm),
	})
}

func (app *application) register(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		form := NewForm(r.PostForm)
		form.required("name", "email", "password", "confirm_password", "avatar").maxLength("email", 50).minLength("password", 6)
		if !form.valid() {
			app.errorLog.Printf("Invalid form data %v", form.Errors)
			form.Errors.Add("generic", "Invalid input")
			app.render(w, r, "register.html", &templateData{
				Form: form,
			})
			return
		}

		email := r.FormValue("email")
		password := r.FormValue("password")
		name := r.FormValue("name")
		avatar := r.FormValue("avatar")
		confirmPassword := r.FormValue("confirm_password")

		if password != confirmPassword {
			form.Errors.Add("generic", "Password and confirm password do not match")
			app.render(w, r, "register.html", &templateData{
				Form: form,
			})
			return
		}

		id, err := app.userRepo.CreateUser(name, email, password, avatar)
		if err != nil {
			form.Errors.Add("generic", err.Error())
			app.render(w, r, "register.html", &templateData{
				Form: form,
			})
			return
		}

		//logged in
		// app.session.Put(r, loggedInUserKey, id)

		app.infoLog.Printf("User registered successfully : %v", id)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	app.render(w, r, "register.html", &templateData{
		Form: NewForm(r.PostForm),
	})
}

func (app *application) logout(w http.ResponseWriter, r *http.Request) {
	app.session.Remove(r, loggedInUserKey)
	app.session.Put(r, "flash", "You logged out successfully !")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
