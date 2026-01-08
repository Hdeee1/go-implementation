package middleware

import (
	"log"
	"net/http"
	"time"
)


func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// log for time
		log.Printf("Started %s %s at %s", r.Method, r.URL.Path, start.Format(time.RFC3339))

		// call handler
		next.ServeHTTP(w, r)

		// count duration 
		duration := time.Since(start)
		log.Printf("Completed %s %s in %v", r.Method, r.URL.Path, duration)
	})
}