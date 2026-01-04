package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/WanKapef/go-api/internal/handler"
	"github.com/WanKapef/go-api/internal/middleware"
	"github.com/WanKapef/go-api/internal/repository"
	"github.com/WanKapef/go-api/internal/service"
	"github.com/gorilla/mux"
)

func TestListUsers(t *testing.T) {
	db := newTestDB(t)
	repo := repository.NewUserRepository(db)
	service := service.NewUserService(repo)
	h := handler.NewUserHandler(service)

	router := mux.NewRouter()
	router.Handle("/users", middleware.ErrorMiddleware(h.Create)).Methods(http.MethodPost)
	router.Handle("/users", middleware.ErrorMiddleware(h.List)).Methods(http.MethodGet)

	// create user
	w1 := httptest.NewRecorder()
	createReq := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader([]byte(`{"name":"t1","email":"t1@test.com"}`)))
	createReq.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w1, createReq)

	if w1.Code != http.StatusCreated {
		t.Fatalf("expected 201, got %d", w1.Code)
	}

	// lista usuários
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	w2 := httptest.NewRecorder()
	router.ServeHTTP(w2, req)

	if w2.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w2.Code)
	}

	var users []map[string]interface{}
	json.Unmarshal(w2.Body.Bytes(), &users)

	if len(users) != 1 {
		t.Errorf("expected 1 user, got %d", len(users))
	}

	if users[0]["email"] != "t1@test.com" {
		t.Fatalf("expected email t1@test.com, got %v", users[0]["email"])
	}
}
