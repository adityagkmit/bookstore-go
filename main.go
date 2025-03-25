package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/adityagkmit/bookstore/routes"

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
	router := chi.NewRouter()

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the Bookstore API!"))
	})

	// Connect to Database
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Setup all routes
	routes.SetupRoutes(router, db)

	fmt.Println("Database connection established: ", db.Name())
	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, router))

}
