package controllers

import (
	"net/http"
)

// LoadingLoginPage loading a login page
func LoadingLoginPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("TESTE"))
	//utils.ExecuteTemplate(w, "login.html", nil)
}
