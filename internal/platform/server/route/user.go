package route

import (
	"github.com/gorilla/mux"
	"github.com/jorgeAM/simple-api/internal/platform/server/handler"
)

func initializeUsersRoutes(r *mux.Router, handler handler.Handler) {
	s := r.PathPrefix("/users").Subrouter()
	s.HandleFunc("", handler.NewUser).Methods("POST")
	s.HandleFunc("", handler.GetUsers).Methods("GET")
	s.HandleFunc("/{id}", handler.GetUser).Methods("GET")
	s.HandleFunc("/{id}", handler.DeleteUser).Methods("DELETE")
}
