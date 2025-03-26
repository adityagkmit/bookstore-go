package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/adityagkmit/bookstore/routes"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

type App struct {
	router *chi.Mux
	db     *mongo.Database
}

// NewApp initializes a new application instance
func NewApp(db *mongo.Database) *App {
	router := chi.NewRouter()
	routes.SetupRoutes(router, db)

	return &App{
		router: router,
		db:     db,
	}
}

// Start runs the HTTP server
func (a *App) Start(ctx context.Context) error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:    ":" + port,
		Handler: a.router,
	}

	// Log server start
	log.Println("Server running on port", port)

	// Start the server in a goroutine
	errChan := make(chan error, 1)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			errChan <- fmt.Errorf("failed to start server: %w", err)
		}
		close(errChan)
	}()

	// Graceful shutdown handling
	select {
	case err := <-errChan:
		return err
	case <-ctx.Done():
		log.Println("Shutting down server...")
		return server.Shutdown(ctx)
	}
}
