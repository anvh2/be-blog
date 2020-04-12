package middlewares

import (
	"net/http"
	"strings"
)

// AuthenHTTPServer authentication request and request URI contain excludeMethods
func AuthenHTTPServer(h http.Handler, authen func(r *http.Request), excludeMethods []string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if authen == nil {
			h.ServeHTTP(w, r)
			return
		}

		for _, exclusion := range excludeMethods {
			if ok := strings.Contains(r.RequestURI, exclusion); ok {
				authen(r)
				break
			}
		}

		h.ServeHTTP(w, r)
	})
}
