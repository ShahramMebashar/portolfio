package server

import "net/http"

func (s *Server) homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Link", "</static/images/shahram.webp>; re=preload; as=image")
	err := s.Render(w, "home", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
