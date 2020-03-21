package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jorgeAM/api/db"
	"github.com/jorgeAM/api/models"
)

// GetUser returns user by id
func GetUser(w http.ResponseWriter, r *http.Request) {
	db := db.GetConnection()
	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	u := models.User{}
	db.Table("Users").First(&u, id)

	if u.ID <= 0 {
		log.Fatalf("user with id %s does not exist", id)
		return
	}

	bytes, err := json.Marshal(u)

	if err != nil {
		log.Fatal("something got wrong to convert to json")
		return
	}

	w.Write(bytes)

}
