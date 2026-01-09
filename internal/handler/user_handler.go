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

// CreateUser cria um usuário
// @Summary Criar usuário
// @Description Cria um novo usuário no banco
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.User true "User info"
// @Success 201 {object} model.User
// @Failure 400 {object} map[string]string
// @Router /users [post]
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

// ListUsers lista usuários com paginação e filtros
// @Summary Listar usuários
// @Description Lista usuários com paginação e filtros
// @Tags users
// @Accept json
// @Produce json
// @Param limit query int false "Limite de resultados"
// @Param offset query int false "Offset para paginação"
// @Param page query int false "Número da página"
// @Param name query string false "Filtrar por nome"
// @Param email query string false "Filtrar por email"
// @Param search query string false "Busca geral"
// @Success 200 {array} model.User
// @Failure 400 {object} map[string]string
// @Router /users [get]
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

// GetByID obtém um usuário por ID
// @Summary Obter usuário por ID
// @Description Obtém um usuário pelo seu ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} model.User
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /users/{id} [get]
func (h *UserHandler) GetByID(w http.ResponseWriter, r *http.Request, id int64) error {
	user, err := h.service.FindByID(id)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(user)
}

// Update atualiza um usuário
// @Summary Atualizar usuário
// @Description Atualiza os dados de um usuário existente
// @Tags users
// @Accept json
// @Produce json
// @Param user body model.User true "User info"
// @Success 200 {object} model.User
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /users [put]
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

// Delete remove um usuário por ID
// @Summary Deletar usuário
// @Description Remove um usuário pelo seu ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 204 "No Content"
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /users/{id} [delete]
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
