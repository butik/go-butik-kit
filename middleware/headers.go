package middleware

import (
	"net/http"
)

func WithHeaders(headers map[string]string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				for k, v := range headers {
					w.Header().Add(k, v)
				}
				next.ServeHTTP(w, r)
			},
		)
	}
}
