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
	{
		URI:                    "/search-users",
		Method:                 http.MethodGet,
		Function:               controllers.LoadUsersPage,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userId}",
		Method:                 http.MethodGet,
		Function:               controllers.LoadUsersProfilePage,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userId}/unfollow",
		Method:                 http.MethodPost,
		Function:               controllers.Unfollow,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/users/{userId}/follow",
		Method:                 http.MethodPost,
		Function:               controllers.Follow,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/profile",
		Method:                 http.MethodGet,
		Function:               controllers.LoadUsersLoggedProfilePage,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/update-user",
		Method:                 http.MethodGet,
		Function:               controllers.LoadUsersUpdatePage,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/update-user",
		Method:                 http.MethodPut,
		Function:               controllers.UpdateUser,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/update-password",
		Method:                 http.MethodGet,
		Function:               controllers.LoadUpdatePasswordPage,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/update-password",
		Method:                 http.MethodPost,
		Function:               controllers.UpdatePassword,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/delete-user",
		Method:                 http.MethodDelete,
		Function:               controllers.DeleteUser,
		RequiresAuthentication: true,
	},
}
