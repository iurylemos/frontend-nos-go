package routes

import (
	"frontend-nos/src/controllers/controller_login"
	"frontend-nos/src/controllers/controller_pages"
	"net/http"
)

var routesLogin = []Route{
	{
		URI:                    "/",
		Method:                 http.MethodGet,
		Function:               controller_pages.LoadScreenLogin,
		RequiredAuthentication: false,
	},
	{
		URI:                    "/login",
		Method:                 http.MethodGet,
		Function:               controller_pages.LoadScreenLogin,
		RequiredAuthentication: false,
	},
	{
		URI:                    "/login",
		Method:                 http.MethodPost,
		Function:               controller_login.Login,
		RequiredAuthentication: false,
	},
}
