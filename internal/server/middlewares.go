package server

import (
	"compress/gzip"
	"fmt"
	"net/http"
	"strings"
)

func gzipHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the client accepts GZIP encoding
		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			w.Header().Set("Content-Encoding", "gzip")

			// Create a pipe to connect the gzip writer and the original writer
			gzipWriter := gzip.NewWriter(w)
			defer gzipWriter.Close()

			// Create a response writer that writes to the gzip writer
			writer := gzipResponseWriter{Writer: gzipWriter, ResponseWriter: w}

			// Serve the request using the wrapped writer
			h.ServeHTTP(writer, r)
		} else {
			// If the client does not accept GZIP, serve normally
			h.ServeHTTP(w, r)
		}
	})
}

// gzipResponseWriter wraps the original ResponseWriter to write compressed data
type gzipResponseWriter struct {
	http.ResponseWriter
	Writer *gzip.Writer
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func cacheControlHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set cache control headers for static assets
		w.Header().Set("Cache-Control", "public, max-age=31536000") // Cache for 1 year
		h.ServeHTTP(w, r)
	})
}

func (s *Server) earlyHintsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		style, _ := s.LoadAsset("css/main.css")
		script, _ := s.LoadAsset("js/main.js")

		w.Header().Add("Link", fmt.Sprintf("</static/%s>; rel=preload; as=style", style))
		w.Header().Add("Link", fmt.Sprintf("</static/%s>; rel=preload; as=script", script))

		w.WriteHeader(http.StatusEarlyHints)

		next.ServeHTTP(w, r)
	})
}
