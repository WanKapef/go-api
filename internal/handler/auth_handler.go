package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/WanKapef/go-api/internal/auth"
	"github.com/WanKapef/go-api/internal/service"

	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	service *service.UserService
}

func NewAuthHandler(s *service.UserService) *AuthHandler {
	return &AuthHandler{service: s}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) error {
	var creds struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		//http.Error(w, "invalid request body", http.StatusBadRequest)
		return err
	}

	user, err := h.service.FindByEmail(creds.Email)
	if err != nil {
		//http.Error(w, "invalid credentials", http.StatusUnauthorized)
		return err
	}

	// compara hash com senha digitada
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		//http.Error(w, "invalid email or password", http.StatusUnauthorized)
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

func (h *AuthHandler) Me(w http.ResponseWriter, r *http.Request) error {
	userIDctx := r.Context().Value("userID")
	if userIDctx == nil {
		//http.Error(w, "User not authenticated", http.StatusUnauthorized)
		return errors.New("user not authenticated")
	}
	userID := userIDctx.(int64)

	user, err := h.service.FindByID(userID)
	if err != nil {
		//http.Error(w, "User not found", http.StatusNotFound)
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(user)
}
