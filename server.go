package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
	"github.com/jorgeAM/api/db"
)

type user struct {
	Username string `json:"username"`
}

func main() {
	db, err := db.GetConnection()
	defer db.Close()

	if err != nil {
		log.Fatal(err)
		return
	}

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
