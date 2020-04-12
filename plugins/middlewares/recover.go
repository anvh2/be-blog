package middlewares

import (
	"log"
	"net/http"
)

// RecoverHTTPServer ...
func RecoverHTTPServer(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Println("[Middleware][recover] recovered", r)
			}
		}()
		h.ServeHTTP(w, r)
	})
}
