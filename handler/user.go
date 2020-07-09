package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jorgeAM/api/models"
	"github.com/jorgeAM/api/service"
	"github.com/jorgeAM/api/utils"
)

// Handler handles all endpoint for user
type Handler struct {
	Service service.UserService
}

// GetUsers retrieve users
func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.Service.Repository.GetUsers()

	if err != nil {
		m := &models.Response{
			Code:    http.StatusInternalServerError,
			Message: "something got wrong to retrieve users",
		}

		utils.DisplayMessage(w, m)
		return
	}

	bytes, err := json.Marshal(&users)

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

// GetUser returns user by id
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	castedID, err := strconv.Atoi(id)

	if err != nil {
		m := &models.Response{
			Code:    http.StatusInternalServerError,
			Message: "something got wrong to retrieve users",
		}

		utils.DisplayMessage(w, m)
		return
	}

	user, err := h.Service.Repository.GetUser(castedID)

	if err != nil || user.ID <= 0 {
		m := &models.Response{
			Code:    http.StatusNotFound,
			Message: "user with id " + id + " does not exist",
		}

		utils.DisplayMessage(w, m)
		return
	}

	bytes, err := json.Marshal(user)

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

// NewUser creates a new user
func (h *Handler) NewUser(w http.ResponseWriter, r *http.Request) {
	u := new(models.User)
	json.NewDecoder(r.Body).Decode(u)

	u, err := h.Service.Repository.NewUser(u)

	if err != nil {
		m := &models.Response{
			Code:    http.StatusBadRequest,
			Message: "something got wrong when try to save user",
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
	w.WriteHeader(http.StatusCreated)
	w.Write(bytes)
}

// UpdateUser updates user by id
func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	castedID, err := strconv.Atoi(id)

	if err != nil {
		m := &models.Response{
			Code:    http.StatusInternalServerError,
			Message: "something got wrong to read id of user users",
		}

		utils.DisplayMessage(w, m)
		return
	}

	u := new(models.User)
	u.ID = castedID
	err = json.NewDecoder(r.Body).Decode(u)

	if err != nil {
		m := &models.Response{
			Code:    http.StatusBadRequest,
			Message: "something got wrong to parsing body",
		}

		utils.DisplayMessage(w, m)
		return
	}

	u, err = h.Service.Repository.UpdateUser(u)

	if err != nil {
		m := &models.Response{
			Code:    http.StatusBadRequest,
			Message: "something got wrong to update user",
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
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

// DeleteUser deletes user by id
func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	castedID, err := strconv.Atoi(id)

	if err != nil {
		m := &models.Response{
			Code:    http.StatusInternalServerError,
			Message: "something got wrong to read id of user users",
		}

		utils.DisplayMessage(w, m)
		return
	}

	err = h.Service.Repository.DeleteUser(castedID)

	if err != nil {
		m := &models.Response{
			Code:    http.StatusNotFound,
			Message: "user with id " + id + " does not exist",
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
