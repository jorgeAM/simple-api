package routes

import (
	"github.com/gorilla/mux"
	"github.com/jorgeAM/api/handler"
)

func initializeUsersRoutes(r *mux.Router, handler handler.Handler) {
	s := r.PathPrefix("/users").Subrouter()
	s.HandleFunc("", handler.NewUser).Methods("POST")
	s.HandleFunc("", handler.GetUsers).Methods("GET")
	s.HandleFunc("/test", handler.Get).Methods("GET")
	s.HandleFunc("/{id}", handler.GetUser).Methods("GET")
	s.HandleFunc("/{id}", handler.UpdateUser).Methods("PUT")
	s.HandleFunc("/{id}", handler.DeleteUser).Methods("DELETE")
}
