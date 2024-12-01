package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

type TemplateCache map[string]*template.Template
type pageData map[string]any

func newTemplateCache() (TemplateCache, error) {
	cache := make(TemplateCache)

	pages, err := filepath.Glob("./assets/templates/*.page.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		t, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		t, err = t.ParseGlob("./assets/templates/*.layout.html")
		if err != nil {
			return nil, err
		}

		name := filepath.Base(page)

		cache[name] = t

	}

	return cache, nil
}

func (app *app) render(w http.ResponseWriter, name string, data pageData) {
	t, ok := app.templateCache[name]

	if !ok {
		http.Error(w, fmt.Sprintf("Cannot load page for %s", name), 500)
		return
	}

	buffer := new(bytes.Buffer)
	err := t.Execute(buffer, data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Cannot execute page for %s", name), 500)
		return
	}

	buffer.WriteTo(w)

}
