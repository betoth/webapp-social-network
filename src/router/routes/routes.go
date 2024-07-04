package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Route represent all routes from web apliccation
type Route struct {
	URI                    string
	Method                 string
	Function               func(w http.ResponseWriter, r *http.Request)
	RequiresAuthentication bool
}

// Configuration add all routers to router
func Configuration(router *mux.Router) *mux.Router {
	fmt.Println("Entrou no ")
	routes := routesLogin

	for _, route := range routes {
		router.HandleFunc(route.URI, route.Function).Methods(route.Method)

	}

	return router
}
