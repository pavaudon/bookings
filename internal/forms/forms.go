package forms

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

// creates a custom form struct, embeds a url.Values object
type Form struct {
	url.Values
	Errors errors
}

// returns true if there are no errors otherwise false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// unitializes a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// checks if form field is in post and not empty
func (f *Form) Has(field string) bool {
	x := f.Get(field)
	if x == "" {
		f.Errors.Add(field, "This field cannot be blank")
		return false
	}
	return true
}

func (f *Form) MinLenght(field string, lengh int) bool {
	x := f.Get(field)
	if len(x) < lengh {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long", lengh))
		return false
	}
	return true
}

func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "invalid email address")
	}
}
