package handler

import (
	"encoding/json"
	"net/http"

	"github.com/WanKapef/go-api/internal/auth"
	"github.com/WanKapef/go-api/internal/service"
)

type AuthHandler struct {
	service *service.UserService
}

func NewAuthHandler(s *service.UserService) *AuthHandler {
	return &AuthHandler{service: s}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) error {
	var creds struct {
		Email string `json:"email"`
	}
	json.NewDecoder(r.Body).Decode(&creds)

	user, err := h.service.FindByEmail(creds.Email)
	if err != nil {
		//http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return err
	}

	token, err := auth.GenerateToken(user.ID)
	if err != nil {
		//http.Error(w, "could not generate token", http.StatusInternalServerError)
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(map[string]string{"token": token})
}
