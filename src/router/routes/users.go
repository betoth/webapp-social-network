package routes

import (
	"net/http"
	"webapp-social-network/src/controllers"
)

var routesUser = []Route{
	{
		URI:                    "/register",
		Method:                 http.MethodGet,
		Function:               controllers.LoadingUserRegisterPage,
		RequiresAuthentication: false,
	},
	{
		URI:                    "/users",
		Method:                 http.MethodPost,
		Function:               controllers.CreateUser,
		RequiresAuthentication: false,
	},
}
