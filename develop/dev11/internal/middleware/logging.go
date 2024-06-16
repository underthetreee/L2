package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

func Log(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		slog.Info("request", "method", r.Method, "uri", r.RequestURI, "time", time.Since(start))
	}
}
