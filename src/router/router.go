package router

import "github.com/gorilla/mux"

// Generate retorn a router with all routers configured
func Generate() *mux.Router {
	return mux.NewRouter()
}
