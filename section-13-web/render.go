package main

import (
	"net/http"
)

func (app *application) render(w http.ResponseWriter, filename string, data interface{}) {
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
	app.tp.Render(w, filename, data)
}
