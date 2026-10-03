package middleware

import (
	"context"
	"expense-api/api/internal/utils"
	"net/http"
)

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
				return
			}
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		tokenStr := cookie.Value
		claim, err := utils.ValidateToken(tokenStr)
		if err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		user := context.WithValue(r.Context(), "token", *claim)
		next.ServeHTTP(w, r.WithContext(user))
	}
}
