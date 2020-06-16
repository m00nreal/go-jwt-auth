package router

import (
	"github.com/gorilla/mux"
	"github.com/m00nreal/go-jwt-auth/api/router/routes"
)

func New() *mux.Router  {
	r := mux.NewRouter().StrictSlash(true)
	return routes.SetupRoutesWithMiddleware(r)
}
