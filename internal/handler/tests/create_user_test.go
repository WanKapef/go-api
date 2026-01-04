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

	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func TestCreateUser(t *testing.T) {
	db := newTestDB(t)

	repo := repository.NewUserRepository(db)
	service := service.NewUserService(repo)
	h := handler.NewUserHandler(service)

	router := mux.NewRouter()
	router.Handle("/users", middleware.ErrorMiddleware(h.Create)).Methods(http.MethodPost)

	payload := []byte(`{"name":"Test","email":"test@example.com"}`)
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader(payload))
	req.Header.Add("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("expected status 201, got %d", w.Code)
	}

	var u map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &u)

	if u["name"] != "Test" {
		t.Errorf("expected name Test, got %v", u["name"])
	}
}
