package forms

type errors map[string][]string

// Add new error message for a given field
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
}

// Gett the first error message
func (e errors) Get(field string) string {
	errStrings := e[field]

	if len(errStrings) == 0 {
		return ""
	}

	return errStrings[0]
}
