package httpx

import "net/http"

type HTTPError struct {
	Status  int    `json:"-"`
	Message string `json:"message"`
}

func (e *HTTPError) Error() string {
	return e.Message
}

func New(status int, message string) *HTTPError {
	return &HTTPError{
		Status:  status,
		Message: message,
	}
}

// Helpers comuns
var (
	ErrNotFound     = New(http.StatusNotFound, "recurso não encontrado")
	ErrBadRequest   = New(http.StatusBadRequest, "requisição inválida")
	ErrUnauthorized = New(http.StatusUnauthorized, "não autorizado")
)
