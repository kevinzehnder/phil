package middlewares

import (
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

// CustomResponseWriter is a struct that implements the http.ResponseWriter and http.Hijacker interfaces.
type CustomResponseWriter struct {
	http.ResponseWriter
	status int
	size   int
}

// WriteHeader is a method of the CustomResponseWriter that stores the HTTP status code.
func (w *CustomResponseWriter) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

// Write is a method of the CustomResponseWriter that stores the size of the response body.
func (w *CustomResponseWriter) Write(b []byte) (int, error) {
	size, err := w.ResponseWriter.Write(b)
	w.size += size
	return size, err
}

// Status is a method of the CustomResponseWriter that returns the stored HTTP status code.
func (w *CustomResponseWriter) Status() int {
	return w.status
}

// Size is a method of the CustomResponseWriter that returns the stored response body size.
func (w *CustomResponseWriter) Size() int {
	return w.size
}

// ZeroLogLogger is a middleware function that logs information about incoming requests using Zerolog.
func ZeroLogGPTLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		clientIP := r.RemoteAddr // You can use a middleware to set the real IP in the request context
		method := r.Method
		path := r.URL.Path

		t := time.Now()
		lw := &CustomResponseWriter{w, http.StatusOK, 0}
		next.ServeHTTP(lw, r)
		latency := time.Since(t)
		status := lw.Status()
		size := lw.Size()

		// don't log 200 responses
		switch {
		case status <= 300:
			return

		case status >= 400:
			log.Info().
				Str("client_ip", clientIP).
				Str("method", method).
				Str("path", path).
				Dur("latency", latency).
				Int("status", status).
				Int("size", size).
				Msg("")
		}
	})
}
