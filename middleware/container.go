package middleware

import (
	"net/http"

	"github.com/sarulabs/di"
)

func WithContainer(app di.Container, logFunc func(msg string)) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return di.HTTPMiddleware(next.ServeHTTP, app, logFunc)
	}
}
