package controllers

import (
	"net/http"
	"webapp-social-network/src/cookie"
)

// Logout make logout of social network
func Logout(w http.ResponseWriter, r *http.Request) {
	cookie.Delete(w)
	http.Redirect(w, r, "/login", 302)

}
