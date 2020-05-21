package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	log.Printf("Server started")
	router := NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Accept-Language", "Content-Type", "Authorization"},
		AllowCredentials: true,
		Debug:            true,
	})

	h := c.Handler(router)

	log.Fatal(http.ListenAndServe(":1984", h))
}
