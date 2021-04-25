package routes

import (
	"frontend-nos/src/controllers/controller_pages"
	"net/http"
)

var routeMain = Route{
	URI:                    "/home",
	Method:                 http.MethodGet,
	Function:               controller_pages.LoadScreenMain,
	RequiredAuthentication: true,
}
