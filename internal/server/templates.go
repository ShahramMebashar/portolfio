package server

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
)

// LoadTemplates loads the templates from the frontend directory and stores them in the server
func (s *Server) LoadTemplates() error {
	views := make(map[string]*template.Template)

	pages, err := filepath.Glob("frontend/pages/*.html")
	if err != nil {
		return err
	}

	partials, _ := filepath.Glob("frontend/partials/*.html")

	functions := template.FuncMap{
		"asset": s.LoadAsset,
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles("frontend/layout/base.html")
		if err != nil {
			return err
		}

		if len(partials) > 0 {
			ts, _ = ts.ParseFiles(partials...)
		}

		t, err := ts.ParseFiles(page)
		if err != nil {
			return err
		}

		// remove .html from the name
		name = name[:len(name)-5]

		views[name] = t
	}

	s.views = views

	return nil
}

// Render renders a template with the given data
func (s *Server) Render(w http.ResponseWriter, name string, data map[string]interface{}) error {
	// Load the templates if we are in dev mode
	// This allows us to see changes without restarting the server
	if s.DevMode {
		s.LoadTemplates()
	}

	tmpl, ok := s.views[name]

	if !ok {
		http.Error(w, "template not found", http.StatusInternalServerError)
		return fmt.Errorf("template not found")
	}

	if data == nil {
		data = make(map[string]interface{})
	}

	data["DevMode"] = s.DevMode

	return tmpl.ExecuteTemplate(w, "base.html", data)
}
