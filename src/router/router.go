package router

import (
	"webapp-social-network/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate retorn a router with all routers configured
func Generate() *mux.Router {
	r := mux.NewRouter()
	return routes.SetUp(r)
}
