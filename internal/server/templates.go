package server

import (
	"html/template"
	"path/filepath"
)

func (s *Server) LoadTemplates() error {
	views := make(map[string]*template.Template)

	pages, err := filepath.Glob("frontend/pages/*.html")

	if err != nil {
		return err
	}

	functions := template.FuncMap{
		"asset": s.LoadAsset,
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).Funcs(functions).ParseFiles("frontend/layout/base.html")
		if err != nil {
			return err
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
