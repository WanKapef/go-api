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

func TestDeleteUser(t *testing.T) {
	db := newTestDB(t)
	repo := repository.NewUserRepository(db)
	service := service.NewUserService(repo)
	h := handler.NewUserHandler(service)

	router := mux.NewRouter()
	router.Handle("/users", middleware.ErrorMiddleware(h.Create)).Methods(http.MethodPost)
	router.Handle("/users/{id}", middleware.ErrorMiddleware(httpx.WithID(h.Delete))).Methods(http.MethodDelete)

	// create
	w1 := httptest.NewRecorder()
	createReq := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader([]byte(`{"name":"Del","email":"del@test.com"}`)))
	createReq.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w1, createReq)

	var u map[string]interface{}
	json.Unmarshal(w1.Body.Bytes(), &u)
	id := int(u["id"].(float64))

	// delete
	delReq := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/users/%d", id), nil)
	delW := httptest.NewRecorder()
	router.ServeHTTP(delW, delReq)

	if delW.Code != http.StatusNoContent {
		t.Errorf("expected 204, got %d", delW.Code)
	}
}
