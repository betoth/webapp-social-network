package routes

import (
	"net/http"
	"webapp-social-network/src/middlewares"

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
	routes = append(routes, routeHome)
	routes = append(routes, routePublications...)
	routes = append(routes, routeLogout)

	for _, route := range routes {
		if route.RequiresAuthentication {
			router.HandleFunc(route.URI,
				middlewares.Logger(middlewares.Authentication(route.Function))).Methods(route.Method)

		} else {
			router.HandleFunc(route.URI,
				middlewares.Logger(route.Function)).Methods(route.Method)
		}

	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
