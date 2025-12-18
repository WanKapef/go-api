package httpx

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type HandlerWithID func(w http.ResponseWriter, r *http.Request, id int64)

func WithID(h HandlerWithID) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		idStr, ok := vars["id"]
		if !ok {
			http.Error(w, "id not provided", http.StatusBadRequest)
			return
		}

		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}

		h(w, r, id)
	}
}
