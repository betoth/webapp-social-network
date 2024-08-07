package utils

import (
	"net/http"
	"text/template"
)

var templates *template.Template

// LoadingTemplates insert html templates in templates variable
func LoadingTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
	templates = template.Must(templates.ParseGlob("views/templates/*.html"))
}

// ExecuteTemplate renderize all templates
func ExecuteTemplate(w http.ResponseWriter, template string, data interface{}) {
	templates.ExecuteTemplate(w, template, data)

}
