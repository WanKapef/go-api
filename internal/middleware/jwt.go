package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"github.com/WanKapef/go-api/internal/auth"
)

type middlewareJWT func(w http.ResponseWriter, r *http.Request) error

func JWT(next middlewareJWT) AppHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			//http.Error(w, "missing auth header", http.StatusUnauthorized)
			return errors.New("missing auth header")
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			//http.Error(w, "invalid auth header", http.StatusUnauthorized)
			return errors.New("invalid auth header")
		}

		claims, err := auth.ValidateToken(parts[1])
		if err != nil {
			//fmt.Print(err.Error())
			//http.Error(w, "invalid token", http.StatusUnauthorized)
			return err
		}

		ctx := context.WithValue(r.Context(), "userID", claims.UserID)
		return next(w, r.WithContext(ctx))
	}
}
