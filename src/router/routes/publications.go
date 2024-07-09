package routes

import (
	"net/http"
	"webapp-social-network/src/controllers"
)

var routePublications = []Route{
	{
		URI:                    "/publications",
		Method:                 http.MethodPost,
		Function:               controllers.CreatePublication,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/publications/{publicationId}/like",
		Method:                 http.MethodPost,
		Function:               controllers.Like,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/publications/{publicationId}/unlike",
		Method:                 http.MethodPost,
		Function:               controllers.Unlike,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/publications/{publicationId}/update",
		Method:                 http.MethodGet,
		Function:               controllers.LoadUpdatePublicationPage,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/publications/{publicationId}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdatePublication,
		RequiresAuthentication: true,
	},
	{
		URI:                    "/publications/{publicationId}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeletePublication,
		RequiresAuthentication: true,
	},
}
