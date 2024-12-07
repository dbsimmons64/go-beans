package forms

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

// Field represents a single form field.
type Field interface {
	Render() string
	GetName() string
	SetValue(value string)
}

// BaseField contains shared properties for all fields.
type Attrs struct {
	Name        string
	Label       string
	Placeholder string
	Required    bool
	Type        string
	Class       string
}

// TextField represents a text input field.
type TextField struct {
	Attrs
	Value string
}

func (t *TextField) Render() string {

	tmpl := `
		<div>
		{{if .Label}}<label for="{{.Name}}">{{.Label}}</label>{{end}}
		<input
			type="text"
			id="{{.Attrs.Name}}"
			name="{{.Attrs.Name}}"
			value="{{.Value}}"
			class="foo bar baz {{.Label}}"
		/>
		</div>
		`

	return renderComponent(tmpl, t)

}

func (t *TextField) GetName() string {
	return t.Name
}

func (t *TextField) SetValue(value string) {
	t.Value = value
}

// SelectField represents a dropdown field.
// type SelectField struct {
// 	Attrs
// 	Options []string
// 	Value   string
// }

// func (s *SelectField) Render() string {
// 	options := ""
// 	for _, option := range s.Options {
// 		selected := ""
// 		if option == s.Value {
// 			selected = " selected"
// 		}
// 		options += fmt.Sprintf(`<option value="%s"%s>%s</option>`, option, selected, option)
// 	}
// 	return fmt.Sprintf(
// 		`<label for="%s">%s</label>
// 		<select id="%s" name="%s">%s</select><br>`,
// 		s.Name, s.Label, s.Name, s.Name, options,
// 	)
// }

// func (s *SelectField) GetName() string {
// 	return s.Name
// }

// func (s *SelectField) SetValue(value string) {
// 	s.Value = value
// }

// CheckboxField represents a checkbox input field.
// type CheckboxField struct {
// 	Attrs
// 	Checked bool
// }

// func (c *CheckboxField) Render() string {
// 	checked := ""
// 	if c.Checked {
// 		checked = " checked"
// 	}
// 	return fmt.Sprintf(
// 		`<label for="%s">%s</label>
// 		<input type="checkbox" id="%s" name="%s"%s><br>`,
// 		c.Name, c.Label, c.Name, c.Name, checked,
// 	)
// }

// func (c *CheckboxField) GetName() string {
// 	return c.Name
// }

// func (c *CheckboxField) SetValue(value string) {
// 	c.Checked = value == "on"
// }

// Form represents the entire form.
// type Form struct {
// 	Fields []Field
// }

func (f *Form) Render() string {
	var sb strings.Builder
	sb.WriteString(`<form method="POST" action="/submit">`)
	for _, field := range f.Fields {
		sb.WriteString(field.Render())
	}
	sb.WriteString(`<button type="submit">Submit</button></form>`)
	return sb.String()
}

func (f *Form) ParseForm(r *http.Request) {
	r.ParseForm()
	for _, field := range f.Fields {
		if value, ok := r.Form[field.GetName()]; ok {
			field.SetValue(value[0])
		}
	}
}

func renderComponent(tmpl string, data any) string {
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

	// return template.HTML(buffer.String())
	return buffer.String()
}

// func label() string {
// 	return `<div><label for="{{.Name}}">{{.Label}}</label>`
// }

func label(name, label string) string {
	if label == "" {
		return ""
	}

	return fmt.Sprintf(`<div><label for="%s">%s</label>`, name, label)
}

func exampleUsage() {
	form := &Form{
		Fields: []Field{
			&TextField{
				Attrs: Attrs{
					Name:        "username",
					Label:       "Username",
					Placeholder: "Enter your username",
					Required:    true,
				},
			},
			// &SelectField{
			// 	Attrs: Attrs{
			// 		Name:  "country",
			// 		Label: "Country",
			// 	},
			// 	Options: []string{"USA", "Canada", "UK"},
			// },
			// &CheckboxField{
			// 	Attrs: Attrs{
			// 		Name:  "subscribe",
			// 		Label: "Subscribe to newsletter",
			// 	},
			// },
		},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintln(w, form.Render())
	})

	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		form.ParseForm(r)
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintln(w, "<h1>Form Submitted</h1>")
		for _, field := range form.Fields {
			fmt.Fprintf(w, "<p>%s: %v</p>", field.GetName(), field)
		}
	})

	http.ListenAndServe(":8080", nil)
}
