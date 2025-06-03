package middleware

import (
	"net/http"

	"github.com/sunnygosdk/go-chi-fullcycle-api/configs"
)

func ContextLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := configs.WithLogger(r.Context())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
