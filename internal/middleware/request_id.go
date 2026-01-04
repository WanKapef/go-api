package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

const RequestIDKey = "request_id"

func RequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// pegar do header se existir, ou gerar novo
		reqID := r.Header.Get("X-Request-ID")
		if reqID == "" {
			reqID = uuid.New().String()
		}

		// colocar no header da resposta
		w.Header().Set("X-Request-ID", reqID)

		// colocar no contexto para handlers poderem pegar
		ctx := context.WithValue(r.Context(), RequestIDKey, reqID)

		// seguir para o próximo handler com o novo contexto
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
