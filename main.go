package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	config "github.com/adityagkmit/bookstore/utils"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	// Initialize router
	r := chi.NewRouter()

	// Connect to Database
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	fmt.Println("Database connection established: ", db.Name())
	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, r))

}
