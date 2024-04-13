package middlewares

import (
	"net/http"
	"time"

	"github.com/nakshatraraghav/notesgen/lib"
)

var log = lib.GetLogger()

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func LogRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wr := &wrappedWriter{w, http.StatusOK}

		next.ServeHTTP(wr, r)

		log.Info().Str("method", r.Method).Str("uri", r.URL.Path).Int("status", wr.statusCode).Str("duration", time.Since(start).String()).Send()
	})
}
