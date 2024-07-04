package routes

import (
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

// SetUp add all routers to router
func SetUp(router *mux.Router) *mux.Router {
	routes := routesLogin
	routes = append(routes, routesUser...)

	for _, route := range routes {
		router.HandleFunc(route.URI, route.Function).Methods(route.Method)

	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
