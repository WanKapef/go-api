package handler_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/WanKapef/go-api/internal/handler"
	"github.com/WanKapef/go-api/internal/httpx"
	"github.com/WanKapef/go-api/internal/middleware"
	"github.com/WanKapef/go-api/internal/repository"
	"github.com/WanKapef/go-api/internal/service"
	"github.com/gorilla/mux"
)

func TestGetUserByID(t *testing.T) {
	db := newTestDB(t)
	repo := repository.NewUserRepository(db)
	service := service.NewUserService(repo)
	h := handler.NewUserHandler(service)

	router := mux.NewRouter()
	router.Handle("/users", middleware.ErrorMiddleware(h.Create)).Methods(http.MethodPost)
	router.Handle("/users/{id}", middleware.ErrorMiddleware(httpx.WithID(h.GetByID))).Methods(http.MethodGet)

	// create user
	w1 := httptest.NewRecorder()
	createReq := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader([]byte(`{"name":"G1","email":"g1@test.com"}`)))
	createReq.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w1, createReq)

	var u map[string]interface{}
	json.Unmarshal(w1.Body.Bytes(), &u)
	id := int(u["id"].(float64))

	// get user by id
	getReq := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/users/%d", id), nil)
	w2 := httptest.NewRecorder()
	router.ServeHTTP(w2, getReq)

	if w2.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w2.Code)
	}
}
