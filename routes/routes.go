package routes

import (
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

// SetupRoutes initializes all application routes
func SetupRoutes(r *chi.Mux, db *mongo.Database) {
	SetupAuthRoutes(r, db) // Authentication Routes
}
