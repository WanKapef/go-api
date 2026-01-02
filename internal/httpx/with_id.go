package httpx

import (
	"net/http"
	"strconv"

	"github.com/WanKapef/go-api/internal/middleware"
	"github.com/gorilla/mux"
)

type HandlerWithID func(w http.ResponseWriter, r *http.Request, id int64) error

func WithID(next HandlerWithID) middleware.AppHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		vars := mux.Vars(r)
		idStr := vars["id"]

		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			return err
		}

		return next(w, r, id)
	}
}
