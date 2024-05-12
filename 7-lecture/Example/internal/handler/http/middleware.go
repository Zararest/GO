package handler

import (
	"log"
	"net/http"
	"time"
)

// LogMiddleware calculate request time.
func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {
			log.Printf("request: %s\n\tstart at: %+v\n\tdur: %fs", r.URL.Path, start, time.Since(start).Seconds())
		}()
		next.ServeHTTP(w, r)
	})
}
