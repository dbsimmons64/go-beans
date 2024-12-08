package forms

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"unicode/utf8"
)

type Form struct {
	Values url.Values
	Errors errors
}

func New(data url.Values) *Form {
	return &Form{
		Values: data,
		Errors: errors(map[string][]string{}),
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


// Ensure the field is a valid currency field.
// Valid in this case means optional digits before the decimal point and 1 or 2 digits
// after the decimal point e.g. 12.34
func (f *Form) ValidAmount(field string) {

	re := regexp.MustCompile(`^\d*(\.\d{1,2})?$`)
	if !re.MatchString(f.Values.Get(field)) {
		f.Errors[field] = append(f.Errors[field], fmt.Sprint("Field must be a valid amount"))
	}
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
Ÿ