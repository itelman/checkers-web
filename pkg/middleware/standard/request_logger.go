package standard

import (
	"net/http"
	"strings"
)

type responseRecorder struct {
	http.ResponseWriter
	statusCode      int
	respSizeInBytes int
}

func (r *responseRecorder) WriteHeader(code int) {
	r.statusCode = code
	r.ResponseWriter.WriteHeader(code)
}

func (r *responseRecorder) Write(b []byte) (int, error) {
	n, err := r.ResponseWriter.Write(b)
	r.respSizeInBytes += n
	return n, err
}

func (m *middleware) RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Detect WebSocket upgrade
		if isWSRequest(r) {
			next.ServeHTTP(w, r)
			return
		}

		rr := &responseRecorder{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(rr, r)

		m.logger.Info("HTTP request",
			"method", r.Method,
			"path", r.URL.Path,
			"proto", r.Proto,
			"status", rr.statusCode,
			"bytes", rr.respSizeInBytes,
		)
	})
}

func isWSRequest(r *http.Request) bool {
	return strings.Contains(strings.ToLower(r.Header.Get("Connection")), "upgrade") &&
		strings.ToLower(r.Header.Get("Upgrade")) == "websocket"
}
