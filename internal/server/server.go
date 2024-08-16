package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

type Server struct {
	views map[string]*template.Template
}

func NewServer() *http.Server {
	// Start the server
	fmt.Println("Server started...")

	server := &Server{}
	err := server.LoadTemplates()

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
