package routes

import (
	"github.com/gorilla/mux"
	"github.com/jorgeAM/api/controllers"
)

func initializeUsersRoutes(r *mux.Router) {
	s := r.PathPrefix("/users").Subrouter()
	s.HandleFunc("/{id}", controllers.GetUser).Methods("GET")
	s.HandleFunc("/{id}", controllers.DeleteUser).Methods("DELETE")
}
