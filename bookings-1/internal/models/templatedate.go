package models

import "github.com/SaiKumarThota9/bookings/internal/form"

// TemplateData holds data sent from ghandlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
	Form      *form.Form
}
