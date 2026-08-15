package main

import (
	"net/http"
)

func (app *application) render(w http.ResponseWriter, r *http.Request, filename string, data *templateData) {
	// filepath := path.Join(app.templateDir, filename)
	// tmp, err := template.ParseFiles(filepath)
	// if err != nil {
	// 	app.errorLog.Println(err)
	// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	// 	return
	// }

	// err = tmp.Execute(w, data)
	// if err != nil {
	// 	app.errorLog.Println(err)
	// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	// 	return
	// }
	app.tp.Render(w, filename, app.defaultTemplateData(data, r))
}

func (app *application) defaultTemplateData(data *templateData, r *http.Request) *templateData {
	if data == nil {
		data = &templateData{}
	}
	data.Flash = app.session.PopString(r, "flash")
	// data.isAuthenticated = a.session.Get(r, "userId") != nil
	return data
}
