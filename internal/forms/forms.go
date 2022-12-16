package forms

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

// Custom form struct to embed url.Values object
type Form struct {
	url.Values
	Errors errors
}

// Initialize the form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Check if a form field is not empty
func (f *Form) Has(field string) bool {
	x := f.Get(field)

	if x == "" {
		return false
	}

	return true
}

// Returns true if there is no errors
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// Validate required fields
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)

		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field is required")
		}
	}
}

// Min length for fields
func (f *Form) MinLength(field string, length int) bool {
	x := f.Get(field)

	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must have at least %d characters", length))
		return false
	}

	return true
}

// Check email field
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid email address")
	}
}
