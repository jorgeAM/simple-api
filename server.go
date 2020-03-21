package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type user struct {
	Username string `json:"username"`
}

func main() {
	os.Setenv("MYSQL_USER", "root")
	eu := os.Getenv("MYSQL_USER")
	fmt.Println("ACA", eu)

	r := mux.NewRouter()
	r.HandleFunc("/user/{id}", getUser).Methods("GET")

	s := http.Server{
		Addr:    ":3000",
		Handler: r,
	}

	log.Println("server is running on port 3000")
	log.Fatal(s.ListenAndServe())
}

func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println(id)
	u := &user{Username: "steezclick"}
	bytes, err := json.Marshal(u)

	if err != nil {
		log.Fatal("something got wrong to convert to json")
		return
	}

	w.Write(bytes)
}
