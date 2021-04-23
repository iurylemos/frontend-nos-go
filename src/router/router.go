package router

import (
	"frontend-nos/src/router/routes"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	return routes.Configure(r)
}
