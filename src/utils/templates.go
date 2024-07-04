package utils

import (
	"fmt"
	"net/http"
	"text/template"
)

var templates *template.Template

// LoadingTemplates insert html templates in templates variable
func LoadingTemplates() {
	fmt.Println("Entrou no LoadingTemplates ")
	templates = template.Must(template.ParseGlob("views/*.html"))
}

// ExecuteTemplate renderize all templates
func ExecuteTemplate(w http.ResponseWriter, template string, data interface{}) {
	fmt.Println("Entrou no ExecuteTemplate ")
	templates.ExecuteTemplate(w, template, data)

}
