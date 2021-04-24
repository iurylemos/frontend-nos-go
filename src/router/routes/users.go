package routes

import (
	"frontend-nos/src/controllers/controller_pages"
	"frontend-nos/src/controllers/controller_users"
	"net/http"
)

var routesUsers = []Route{
	{
		URI:                    "/criar-usuario",
		Method:                 http.MethodGet,
		Function:               controller_pages.LoadScreenRegister,
		RequiredAuthentication: false,
	},
	{
		URI:                    "/usuarios",
		Method:                 http.MethodPost,
		Function:               controller_users.RegisterUser,
		RequiredAuthentication: false,
	},
}
