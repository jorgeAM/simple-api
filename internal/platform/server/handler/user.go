package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jorgeAM/simple-api/internal/kit/response"
	"github.com/jorgeAM/simple-api/internal/user/application/creating"
	"github.com/jorgeAM/simple-api/internal/user/application/removing"
	"github.com/jorgeAM/simple-api/internal/user/application/retrieve"
	"github.com/jorgeAM/simple-api/internal/user/application/updating"
	"github.com/jorgeAM/simple-api/internal/user/domain"
)

// Handler handles all endpoint for user
type Handler struct {
	Creating   *creating.UserCreatingService
	Retrieving *retrieve.UserRetrieveAllService
	Finding    *retrieve.UserRetrieveOneService
	Updating   *updating.UserUpdatingService
	Removing   *removing.UserRemovingService
}

type createUserRequest struct {
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

	if err != nil || user.ID == "" {
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
			Message: "something got wrong when try to save user",
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

func (h *Handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	u := new(domain.User)
	u.ID = id
	err := json.NewDecoder(r.Body).Decode(u)

	if err != nil {
		m := &response.Response{
			Code:    http.StatusBadRequest,
			Message: "something got wrong to parsing body",
		}

		response.DisplayMessage(w, m)
		return
	}

	u, err = h.Updating.UpdateUser(context.Background(), u)

	if err != nil {
		m := &response.Response{
			Code:    http.StatusBadRequest,
			Message: "something got wrong to update user",
		}

		response.DisplayMessage(w, m)
		return
	}

	bytes, err := json.Marshal(u)

	if err != nil {
		m := &response.Response{
			Code:    http.StatusBadRequest,
			Message: "something got wrong to convert to json",
		}

		response.DisplayMessage(w, m)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
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
