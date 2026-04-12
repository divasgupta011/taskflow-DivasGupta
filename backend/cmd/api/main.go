package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"taskflow/internal/app"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	application := app.New()

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      application.Router(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Printf("server running on port %s", port)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
