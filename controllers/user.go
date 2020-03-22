package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jorgeAM/api/db"
	"github.com/jorgeAM/api/models"
	"github.com/jorgeAM/api/utils"
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
		m := &models.Response{
			Code:    http.StatusNotFound,
			Message: "user with id " + id + " does not exist",
		}

		utils.DisplayMessage(w, m)
		return
	}

	bytes, err := json.Marshal(u)

	if err != nil {
		m := &models.Response{
			Code:    http.StatusBadRequest,
			Message: "something got wrong to convert to json",
		}

		utils.DisplayMessage(w, m)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

// DeleteUser delete user by id
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db := db.GetConnection()
	defer db.Close()

	vars := mux.Vars(r)
	id := vars["id"]

	u := models.User{}
	db.Table("Users").First(&u, id)

	if u.ID <= 0 {
		m := &models.Response{
			Code:    http.StatusNotFound,
			Message: "user with id " + id + " does not exist",
		}

		utils.DisplayMessage(w, m)
		return
	}

	err := db.Table("Users").Delete(&u).Error

	if err != nil {
		m := &models.Response{
			Code:    http.StatusBadRequest,
			Message: "something got wrong when try to delete record",
		}

		utils.DisplayMessage(w, m)
		return
	}

	m := &models.Response{
		Code:    http.StatusOK,
		Message: "user was deleted successfully",
	}

	utils.DisplayMessage(w, m)
}
