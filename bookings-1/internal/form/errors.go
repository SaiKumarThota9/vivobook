package form

type errors map[string][]string

// Adds an error message for a given form field
func (e errors) Add(field string, message string) {
	e[field] = append(e[field], message)
}
//Get returns the first error message
func (e errors) Get(field string) string {
	es := e[field]
	if len(field) == 0 {
		return ""
	}
	return es[0]
}
