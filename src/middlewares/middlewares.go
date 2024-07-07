package middlewares

import (
	"log"
	"net/http"
	"webapp-social-network/src/cookie"
)

// Logger print requisitions information in terminal
func Logger(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunc(w, r)
	}
}

// Authentication verify cookies existence
func Authentication(nextFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := cookie.Read(r)

		if err != nil {
			http.Redirect(w, r, "/login", 302)
			return
		}

		nextFunc(w, r)
	}
}
