package main

import (
	"log"
	"net/http"
	"calc_service/internal/handlers"
	"calc_service/internal/middleware"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/calculate", middleware.ValidationMiddleware(handlers.CalculateHandler))

	log.Println("Starting server on :8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Server failed: %s", err)
	}
}
