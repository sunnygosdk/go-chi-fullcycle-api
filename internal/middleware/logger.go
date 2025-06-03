package middleware

import (
	"net/http"

	"github.com/sunnygosdk/go-chi-fullcycle-api/config"
)

func ContextLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := config.WithLogger(r.Context())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
