package http

import "net/http"

func JWTAuth(
	original func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header["Authorization"]
		if authHeader == nil {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}
	}
}
