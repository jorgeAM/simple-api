package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jorgeAM/simple-api/internal/kit/response"
	"github.com/jorgeAM/simple-api/internal/user/application/creating"
	"github.com/jorgeAM/simple-api/internal/user/application/finding"
	"github.com/jorgeAM/simple-api/internal/user/application/removing"
	"github.com/jorgeAM/simple-api/internal/user/application/retrieve"
)

// Handler handles all endpoint for user
type Handler struct {
	Creating   *creating.UserCreatingService
	Retrieving *retrieve.UserRetrieveAllService
	Finding    *finding.UserRetrieveOneService
	Removing   *removing.UserRemovingService
}

type createUserRequest struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type updateUserRequest struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func (h *Handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.Retrieving.GetAllUser(r.Context())

	if err != nil {
		m := &response.Response{
			Code:    http.StatusInternalServerError,
			Message: "something got wrong to retrieve users",
		}

		response.DisplayMessage(w, m)
		return
	}

	bytes, err := json.Marshal(users)

	if err != nil {
		m := &response.Response{
			Code:    http.StatusBadRequest,
			Message: "something got wrong to convert to json",
		}

		response.DisplayMessage(w, m)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	user, err := h.Finding.FindUserByID(context.Background(), id)

	if err != nil {
		m := &response.Response{
			Code:    http.StatusNotFound,
			Message: "user with id " + id + " does not exist",
		}

		response.DisplayMessage(w, m)
		return
	}

	bytes, err := json.Marshal(user)

	if err != nil {
		m := &response.Response{
			Code:    http.StatusBadRequest,
			Message: "something got wrong to convert to json",
		}

		response.DisplayMessage(w, m)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(bytes)
}

func (h *Handler) NewUser(w http.ResponseWriter, r *http.Request) {
	var req createUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		m := &response.Response{
			Code:    http.StatusBadRequest,
			Message: "something got wrong to parse request body",
		}

		response.DisplayMessage(w, m)
		return
	}

	err = h.Creating.CreateNewUser(context.Background(), req.ID, req.Username, req.FirstName, req.LastName)

	if err != nil {
		m := &response.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}

		response.DisplayMessage(w, m)
		return
	}

	m := &response.Response{
		Code:    http.StatusCreated,
		Message: "user created successfully",
	}

	response.DisplayMessage(w, m)
}

func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	err := h.Removing.RemoveUserByID(context.Background(), id)

	if err != nil {
		m := &response.Response{
			Code:    http.StatusNotFound,
			Message: "user with id " + id + " does not exist",
		}

		response.DisplayMessage(w, m)
		return
	}

	m := &response.Response{
		Code:    http.StatusOK,
		Message: "user was deleted successfully",
	}

	response.DisplayMessage(w, m)
}
