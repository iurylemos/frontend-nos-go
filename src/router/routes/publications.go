package routes

import (
	"frontend-nos/src/controllers/controller_publications"
	"net/http"
)

var routesPublications = []Route{
	{
		URI:                    "/publicacoes",
		Method:                 http.MethodPost,
		Function:               controller_publications.CreatePublication,
		RequiredAuthentication: true,
	},
}
