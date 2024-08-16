package main

import (
	"log"

	"github.com/ShahramMebashar/portolfio/internal/server"
)

func main() {
	server := server.NewServer()

	log.Println("Server started at http://localhost:8000")

	err := server.ListenAndServe()

	if err != nil {
		log.Fatalf("server failed with %s", err)
	}
}
