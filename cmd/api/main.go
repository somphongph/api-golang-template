package main

import (
	"api-places/internal/server"
	"log"
)

func main() {
	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("cannot start server: %s", err)
	}
}
