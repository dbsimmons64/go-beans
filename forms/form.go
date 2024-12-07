package forms

import (
	"fmt"
	"net/url"
	"strings"
	"unicode/utf8"
)

type Form struct {
	Values url.Values
	Errors errors
	Fields []Field
}

func New(data url.Values) *Form {
	return &Form{
		Values: data,
		Errors: errors(map[string][]string{}),
		Fields: []Field{},
	}
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		if strings.TrimSpace(f.Values.Get(field)) == "" {
			f.Errors[field] = append(f.Errors[field], "This field cannot be empty")
		}
	}
}

func (f *Form) MinLength(field string, length int) {
	if utf8.RuneCountInString(f.Values.Get(field)) < length {
		f.Errors[field] = append(f.Errors[field], fmt.Sprintf("This field must be at least %d characters", length))
	}
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
