package routes

import (
	"github.com/adityagkmit/bookstore/handlers"
	"github.com/adityagkmit/bookstore/repositories"
	"github.com/adityagkmit/bookstore/services"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupAuthRoutes(r *chi.Mux, db *mongo.Database) {
	authRepo := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(authRepo)
	authHandler := handlers.NewAuthHandler(authService)

	r.Post("/api/register", authHandler.Register)
	r.Post("/api/login", authHandler.Login)
}
