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

func TestUpdateUser(t *testing.T) {
	db := newTestDB(t)
	repo := repository.NewUserRepository(db)
	service := service.NewUserService(repo)
	h := handler.NewUserHandler(service)

	router := mux.NewRouter()
	router.Handle("/users", middleware.ErrorMiddleware(h.Create)).Methods(http.MethodPost)
	router.Handle("/users", middleware.ErrorMiddleware(h.Update)).Methods(http.MethodPut)

	// create user
	w1 := httptest.NewRecorder()
	createReq := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader([]byte(`{"name":"Old","email":"old@test.com"}`)))
	createReq.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w1, createReq)

	var u map[string]interface{}
	json.Unmarshal(w1.Body.Bytes(), &u)
	u["name"] = "Updated"

	// update
	payload, _ := json.Marshal(u)
	updateReq := httptest.NewRequest(http.MethodPut, "/users", bytes.NewReader(payload))
	updateReq.Header.Add("Content-Type", "application/json")
	w2 := httptest.NewRecorder()
	router.ServeHTTP(w2, updateReq)

	if w2.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w2.Code)
	}
}
