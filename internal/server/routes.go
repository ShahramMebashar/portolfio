package server

import (
	"net/http"
)

func (s *Server) RegisterRoutes() http.Handler {
	// Register the routes
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("public"))

	mux.Handle("/static/", gzipHandler(cacheControlHandler(http.StripPrefix("/static/", fs))))
	mux.Handle("/robots.txt", gzipHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "public/robots.txt")
	})))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := s.Render(w, "home", nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	return s.earlyHintsMiddleware(mux)
}
