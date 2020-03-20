package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/user", getUser).Methods("GET")

	s := http.Server{
		Addr:    ":3000",
		Handler: r,
	}

	log.Fatal(s.ListenAndServe())
}

func getUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hola!"))
}
