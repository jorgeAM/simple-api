package routes

import "github.com/gorilla/mux"

// InitializeRoutes initialize all endpoints
func InitializeRoutes() *mux.Router {
	r := mux.NewRouter()
	initializeUsersRoutes(r)
	return r
}
