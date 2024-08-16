package server

import (
	"net/http"
)

func (s *Server) RegisterRoutes() http.Handler {
	// Register the routes
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("public"))

	mux.Handle("/static/", http.StripPrefix("/static/", fs))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, ok := s.views["home"]

		if !ok {
			http.Error(w, "template not found", http.StatusInternalServerError)
			return
		}

		err := tmpl.ExecuteTemplate(w, "base.html", map[string]any{
			"DevMode": true,
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	return mux
}
