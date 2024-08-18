package server

import (
	"compress/gzip"
	"fmt"
	"net/http"
	"strings"
)

type gzipResponseWriter struct {
	http.ResponseWriter
	Writer *gzip.Writer
}

func (g *gzipResponseWriter) Write(p []byte) (int, error) {
	return g.Writer.Write(p)
}

func gzipMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			w.Header().Add("Content-Encoding", "gzip")

			gzipWriter := gzip.NewWriter(w)
			defer gzipWriter.Close()

			gzipResponseWriter := &gzipResponseWriter{ResponseWriter: w, Writer: gzipWriter}

			next.ServeHTTP(gzipResponseWriter, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func cacheControlMiddleware(h http.Handler) http.Handler {
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
