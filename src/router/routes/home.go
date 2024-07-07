package routes

import (
	"net/http"
	"webapp-social-network/src/controllers"
)

var routeHome = Route{
	URI:                    "/home",
	Method:                 http.MethodGet,
	Function:               controllers.LoadingHomePage,
	RequiresAuthentication: true,
}
