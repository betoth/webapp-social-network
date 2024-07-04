package controllers

import (
	"net/http"
	"webapp-social-network/src/utils"
)

// LoadingLoginPage loading a login page
func LoadingLoginPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "login.html", nil)
}

// LoadingUserRegisterPage loading a users register page
func LoadingUserRegisterPage(w http.ResponseWriter, r *http.Request) {
	utils.ExecuteTemplate(w, "register.html", r)

}
