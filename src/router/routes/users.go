package routes

import (
	"frontend-nos/src/controllers/controller_pages"
	"net/http"
)

var routesUsers = []Route{
	{
		URI:                    "/criar-usuario",
		Method:                 http.MethodGet,
		Function:               controller_pages.LoadScreenRegister,
		RequiredAuthentication: false,
	},
}
