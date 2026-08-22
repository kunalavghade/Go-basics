package main

import (
	"fmt"
	"net/http"
	"path"
	"path/filepath"
	"sync"
	"text/template"
)

type TemplateRenderer struct {
	cache       map[string]*template.Template
	mutex       sync.Mutex
	dev         bool
	templateDir string
}

type templateData struct {
	Form            *Form
	IsAuthenticated bool
	Flash           string
}

func NewTemplateRenderer(isDev bool, dir string) *TemplateRenderer {
	return &TemplateRenderer{
		cache:       make(map[string]*template.Template),
		mutex:       sync.Mutex{},
		dev:         isDev,
		templateDir: dir,
	}
}

func (t *TemplateRenderer) Render(w http.ResponseWriter, name string, data interface{}) {
	tmpl, err := t.getTemplate(name)
	if err != nil {
		fmt.Println("getTemplate error:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteTemplate(w, "base.html", data)
	if err != nil {
		fmt.Println("Template error:", err)
		http.Error(w, fmt.Sprintf("Internal Server Error : %v", err), http.StatusInternalServerError)
		return
	}
}

func (t *TemplateRenderer) getTemplate(name string) (*template.Template, error) {
	if !t.dev {
		t.mutex.Lock()
		if tmpl, ok := t.cache[name]; ok {
			t.mutex.Unlock()
			return tmpl, nil
		}
		t.mutex.Unlock()
	}
	tmpl, err := t.templete(name)
	if err != nil {
		return nil, err
	}
	if !t.dev {
		t.mutex.Lock()
		t.cache[name] = tmpl
		t.mutex.Unlock()
	}
	return tmpl, nil
}

func (t *TemplateRenderer) templete(name string) (*template.Template, error) {
	tmpletePath := path.Join(t.templateDir, name)
	files := []string{tmpletePath}

	layoutPath := path.Join(t.templateDir, "layouts/*.html")
	layout, err := filepath.Glob(layoutPath)
	if err == nil {
		files = append(files, layout...)
	}

	partialsPath := path.Join(t.templateDir, "partials/*.html")
	partials, err := filepath.Glob(partialsPath)
	if err == nil {
		files = append(files, partials...)
	}

	return template.ParseFiles(files...)
}
