package server

import (
	"html/template"
	"path/filepath"
)

func (s *Server) LoadTemplates() error {
	views := make(map[string]*template.Template)

	layouts, err := filepath.Glob("web/layout/*.html")

	if err != nil {
		return err
	}

	pages, err := filepath.Glob("web/pages/*.html")

	if err != nil {
		return err
	}

	for _, page := range pages {
		files := append(layouts, page)
		name := filepath.Base(page)

		t, err := template.ParseFiles(files...)

		if err != nil {
			return err
		}

		// name are like home.html
		// we want to remove the .html part
		name = name[:len(name)-5]
		views[name] = t
	}

	s.views = views

	return nil
}
