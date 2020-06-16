package routes

import (
	"github.com/m00nreal/go-jwt-auth/api/controllers"
	"net/http"
)

var usersRoutes = []Route {
	Route {
		Uri: "/users",
		Method: http.MethodGet,
		Handler: controllers.GetUsers,
	},
	Route {
		Uri: "/users/{id}",
		Method: http.MethodGet,
		Handler: controllers.GetUser,
	},
	Route {
		Uri: "/users",
		Method: http.MethodPost,
		Handler: controllers.CreateUser,
	},
	Route {
		Uri: "/users/{id}",
		Method: http.MethodDelete,
		Handler: controllers.DeleteUser,
	},
	Route {
		Uri: "/users/{id}",
		Method: http.MethodPut,
		Handler: controllers.UpdateUser,
	},
}
