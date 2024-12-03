package forms

import (
	"fmt"
	"net/url"
	"strings"
	"unicode/utf8"
)

type Form struct {
	url.Values
	Errors errors
}

func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		if strings.TrimSpace(f.Get(field)) == "" {
			f.Errors[field] = append(f.Errors[field], "This field cannot be empty")
		}
	}
}

func (f *Form) MinLength(field string, length int) {
	if utf8.RuneCountInString(f.Get(field)) < length {
		f.Errors[field] = append(f.Errors[field], fmt.Sprintf("This field must be at least %d characters", length))
	}
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
