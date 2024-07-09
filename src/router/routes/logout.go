package routes

import (
	"net/http"
	"webapp-social-network/src/controllers"
)

var routeLogout = Route{
	URI:                    "/logout",
	Method:                 http.MethodGet,
	Function:               controllers.Logout,
	RequiresAuthentication: true,
}
