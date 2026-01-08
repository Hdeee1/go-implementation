package middleware

import "net/http"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := r.Header.Get("Authorization"); err == "" {
			http.Error(w, "Unauthorized", 401)
			return 
		}
		next.ServeHTTP(w, r)
	})
	
}