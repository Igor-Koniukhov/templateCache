package models

// TemplatedData - data that could possibly be put in the template
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string // CSRFToken - cross site request forgery token
	Flash     string
	Warning   string
	Error     string
}