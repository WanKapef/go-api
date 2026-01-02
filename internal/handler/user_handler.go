package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/WanKapef/go-api/internal/model"
	"github.com/WanKapef/go-api/internal/service"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(s *service.UserService) *UserHandler {
	return &UserHandler{service: s}
}

// Create
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) error {
	var user model.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return err
	}

	if err := h.service.CreateUser(&user); err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(user)
}

// Read
func (h *UserHandler) List(w http.ResponseWriter, r *http.Request) error {
	query := r.URL.Query()

	limit, _ := strconv.Atoi(query.Get("limit"))
	offset, _ := strconv.Atoi(query.Get("offset"))
	page, _ := strconv.Atoi(query.Get("page"))

	name := query.Get("name")
	email := query.Get("email")
	search := query.Get("search")

	users, err := h.service.ListUsers(limit, offset, page, name, email, search)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) ListByID(w http.ResponseWriter, r *http.Request, id int64) error {
	user, err := h.service.ListByID(id)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(user)
}

// Update
func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) error {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		return err
	}

	if err := h.service.UpdateUser(&user); err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(user)
}

// Delete
func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request, id int64) error {
	if id <= 0 {
		return errors.New("ID inválido")
	}

	if err := h.service.DeleteUser(id); err != nil {
		return err
	}

	w.WriteHeader(http.StatusNoContent)
	return nil
}
