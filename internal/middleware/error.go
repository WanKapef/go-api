package middleware

import (
	"encoding/json"
	"net/http"
)

type AppHandler func(w http.ResponseWriter, r *http.Request) error

func ErrorMiddleware(next AppHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := next(w, r)

		if err == nil {
			return
		}

		// resposta padronizada JSON
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(map[string]any{
			"error": err.Error(),
		})
	})
}
