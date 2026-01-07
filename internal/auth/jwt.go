package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/WanKapef/go-api/internal/config"

	"github.com/golang-jwt/jwt/v5"
)

var JWTSecret = config.Load().JWTSecret

// claims customizadas
type Claims struct {
	UserID int64 `json:"user_id"`
	jwt.RegisteredClaims
}

// gera token JWT
func GenerateToken(userID int64) (string, error) {
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(JWTSecret))
}

// valida token
func ValidateToken(tokenStr string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWTSecret), nil
	})
	if err != nil {
		// erros específicos da v5
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, fmt.Errorf("token expirado")
		}
		if errors.Is(err, jwt.ErrTokenSignatureInvalid) {
			return nil, fmt.Errorf("assinatura inválida")
		}
		return nil, fmt.Errorf("token inválido: %w", err)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("token inválido")
	}

	return claims, nil
}
