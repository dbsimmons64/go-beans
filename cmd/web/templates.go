package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/dbsimmons64/go-beans/forms"
)

type TemplateCache map[string]*template.Template
type pageData map[string]any

func newTemplateCache() (TemplateCache, error) {
	cache := make(TemplateCache)

	pages, err := filepath.Glob("./assets/templates/02-pages/*.page.html")
	if err != nil {
		return nil, err
	}

	for _, page := range pages {

		t, err := template.New(page).Funcs(template.FuncMap{
			"inputField": inputField,
			"dateField":  dateField,
		}).ParseFiles(page)

		// t, err := t.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		t, err = t.ParseGlob("./assets/templates/01-Blocks/*.layout.html")
		if err != nil {
			return nil, err
		}

		t, err = t.ParseGlob("./assets/templates/03-Layouts/*.layout.html")

		if err != nil {
			return nil, err
		}

		name := filepath.Base(page)

		cache[name] = t

	}

	return cache, nil
}

func (app *app) renderPage(w http.ResponseWriter, name string, data pageData) {
	t, ok := app.templateCache[name]

	if !ok {
		http.Error(w, fmt.Sprintf("Cannot load page for %s", name), 500)
		return
	}

	buffer := new(bytes.Buffer)
	err := t.ExecuteTemplate(buffer, name, data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Cannot execute page %s, error: %s", name, err), 500)
		return
	}

	buffer.WriteTo(w)

}

func renderComponent(tmpl string, data any) template.HTML {
	// Parse and execute the template
	t, err := template.New("fieldTemplate").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	// Execute the template and return the result as a template.HTML

	var buffer bytes.Buffer
	err = t.Execute(&buffer, data)
	if err != nil {
		panic(err)
	}

	return template.HTML(buffer.String())
}

func inputField(form *forms.Form, field, label string) template.HTML {

	data := struct {
		Field  string
		Value  string
		Label  string
		Errors []string
	}{
		Field:  field,
		Value:  form.Values.Get(field),
		Label:  label,
		Errors: form.Errors.Get(field),
	}

	tmpl := `
		<div>
			<label for="{{.Field}}" class="block text-sm font-medium text-gray-700">
				{{.Label}} 
			</label>
			<input 
				type="text" 
				id="{{.Field}}" 
				name="{{.Field}}" 
				value='{{.Value}}' 
				class="pl-2 h-6 mt-1 block w-full border border-black-900 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
			/>
			{{range .Errors}}
				<p class="text-red-600 mt-2 text-sm">{{.}}</p>
			{{end}}
		</div>
	`
	return renderComponent(tmpl, data)

}

func dateField(form *forms.Form, field, label string) template.HTML {

	data := struct {
		Field  string
		Value  string
		Label  string
		Errors []string
	}{
		Field:  field,
		Value:  form.Values.Get(field),
		Label:  label,
		Errors: form.Errors.Get(field),
	}

	tmpl := `
		<div>
			<label for="{{.Field}}" class="block text-sm font-medium text-gray-700">
				{{.Label}} 
			</label>
			<input 
				type="date" 
				id="{{.Field}}" 
				name="{{.Field}}" 
				value='{{.Value}}' 
				class="pl-2 h-6 mt-1 block w-full border border-black-900 rounded-md shadow-sm focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm"
			/>
			{{range .Errors}}
				<p class="text-red-600 mt-2 text-sm">{{.}}</p>
			{{end}}
		</div>
	`
	return renderComponent(tmpl, data)

}
