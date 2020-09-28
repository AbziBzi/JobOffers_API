package middlewares

import (
	"errors"
	"net/http"

	"github.com/abzibzi/jobOfferts_API/api/auth"
	"github.com/abzibzi/jobOfferts_API/api/responses"
)

// SetMiddlewareJSON sets content-type to json
func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

// SetMiddlewareAuthentication verify token and add userID to the request context
func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := auth.TokenValid(r)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
			return
		}
		next(w, r)
	}
}
