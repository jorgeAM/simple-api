package routes

import (
	"github.com/gorilla/mux"
	"github.com/jorgeAM/api/handler"
	"github.com/jorgeAM/api/service"
)

// InitializeRoutes initialize all endpoints
func InitializeRoutes(service service.UserService) *mux.Router {
	r, handler := mux.NewRouter(), handler.Handler{Service: service}
	initializeUsersRoutes(r, handler)
	return r
}
