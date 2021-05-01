package routes

import (
	"frontend-nos/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	URI                    string
	Method                 string
	Function               func(http.ResponseWriter, *http.Request)
	RequiredAuthentication bool
}

func Configure(router *mux.Router) *mux.Router {
	routes := routesLogin

	routes = append(routes, routesUsers...)

	routes = append(routes, routeMain)

	routes = append(routes, routesPublications...)

	for _, route := range routes {

		if route.RequiredAuthentication {
			router.HandleFunc(route.URI,
				middlewares.Logger(
					middlewares.Autheticate(route.Function))).
				Methods(route.Method)
		} else {
			router.HandleFunc(route.URI,
				middlewares.Logger(route.Function)).
				Methods(route.Method)
		}

	}

	// configure folder "assets" but could there is more folders, could be well more
	fileServer := http.FileServer(http.Dir("./assets/"))

	// configure prefix to that doesn't have that go back folders using ".." no html
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
