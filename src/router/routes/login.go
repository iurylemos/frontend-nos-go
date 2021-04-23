package routes

import (
	controller_login "frontend-nos/src/controllers/controller-login"
	"net/http"
)

var routesLogin = []Route{
	{
		URI:                    "/",
		Method:                 http.MethodGet,
		Function:               controller_login.LoadScreenLogin,
		RequiredAuthentication: false,
	},
	{
		URI:                    "/login",
		Method:                 http.MethodGet,
		Function:               controller_login.LoadScreenLogin,
		RequiredAuthentication: false,
	},
}
