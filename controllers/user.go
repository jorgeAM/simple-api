package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jorgeAM/api/models"
)

// GetUser returns user by id
func GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println(id)

	u := models.User{}
	bytes, err := json.Marshal(u)

	if err != nil {
		log.Fatal("something got wrong to convert to json")
		return
	}

	w.Write(bytes)

}
