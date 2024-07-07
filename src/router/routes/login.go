package routes

import (
	"net/http"
	"webapp-social-network/src/controllers"
)

var routesLogin = []Route{
	{
		URI:                    "/",
		Method:                 http.MethodGet,
		Function:               controllers.LoadingLoginPage,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/login",
		Method:                 http.MethodGet,
		Function:               controllers.LoadingLoginPage,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/login",
		Method:                 http.MethodPost,
		Function:               controllers.Login,
		RequiresAuthentication: false,
	},
}
