package main

import (
	"fmt"
	"go-warehousing/configuration"
	"go-warehousing/controller"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	var db configuration.DB
	// Initialize the database
	db.Connect()

	mux := http.NewServeMux()
	srv := configuration.NewServer(":8080", mux)
	srv.Run()

	// handler
	mux.HandleFunc("/", controller.HandlerHomepage)
	fmt.Println("Server is listening on port 8080")

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 5 seconds.
	// In a real-world application, you might use a signal handler for graceful shutdown.
	select {}

}
