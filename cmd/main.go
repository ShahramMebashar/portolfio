package main

import (
	"fmt"
	"log"

	"github.com/ShahramMebashar/portolfio/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("failed to load env with %s", err)
	}

	server := server.NewServer()

	fmt.Println("Server started at http://localhost:8000")

	err = server.ListenAndServe()

	if err != nil {
		log.Fatalf("server failed with %s", err)
	}
}
