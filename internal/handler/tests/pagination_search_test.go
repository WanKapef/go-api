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

func TestPaginationAndSearch(t *testing.T) {
	db := newTestDB(t)
	repo := repository.NewUserRepository(db)
	service := service.NewUserService(repo)
	h := handler.NewUserHandler(service)

	router := mux.NewRouter()
	router.Handle("/users", middleware.ErrorMiddleware(h.Create)).Methods(http.MethodPost)
	router.Handle("/users", middleware.ErrorMiddleware(h.List)).Methods(http.MethodGet)

	// add users
	users := []string{
		`{"name":"Alice","email":"alice@test.com"}`,
		`{"name":"Bob","email":"bob@test.com"}`,
		`{"name":"Alex","email":"alex@test.com"}`,
	}

	for _, u := range users {
		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewReader([]byte(u)))
		req.Header.Add("Content-Type", "application/json")
		httptest.NewRecorder()
		router.ServeHTTP(httptest.NewRecorder(), req)
	}

	// test pagination
	reqPag := httptest.NewRequest(http.MethodGet, "/users?limit=2&offset=1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, reqPag)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200, got %d", w.Code)
	}

	var res []map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &res)

	if len(res) != 2 {
		t.Errorf("expected 2, got %d", len(res))
	}

	// test search by name
	reqSearch := httptest.NewRequest(http.MethodGet, "/users?name=Al", nil)
	w2 := httptest.NewRecorder()
	router.ServeHTTP(w2, reqSearch)

	var resSearch []map[string]interface{}
	json.Unmarshal(w2.Body.Bytes(), &resSearch)

	if len(resSearch) == 0 {
		t.Errorf("expected at least 1 user in search, got %d", len(resSearch))
	}
}
