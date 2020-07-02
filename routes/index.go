package routes

import (
	"github.com/gorilla/mux"
	"github.com/jorgeAM/api/controllers"
	"github.com/jorgeAM/api/repository"
)

// InitializeRoutes initialize all endpoints
func InitializeRoutes(repository repository.Repository) *mux.Router {
	r, handler := mux.NewRouter(), controllers.Handler{Repository: repository}
	initializeUsersRoutes(r, handler)
	return r
}
