package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ShahramMebashar/portolfio/internal/vite"
)

type Server struct {
	views    map[string]*template.Template
	manifest map[string]*vite.ManifestEntry
	DevMode  bool
}

func NewServer() *http.Server {
	manifest, err := vite.LoadViteManifest()

	if err != nil {
		log.Fatalf("failed to load vite manifest with %s", err)
	}

	appMode := os.Getenv("APP_MODE")

	server := &Server{
		manifest: manifest,
		DevMode:  appMode != "production",
	}

	err = server.LoadTemplates()

	if err != nil {
		log.Fatalf("failed to load templates with %s", err)
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	return &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      server.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}
