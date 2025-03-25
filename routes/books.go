package routes

import (
	"github.com/adityagkmit/bookstore/handlers"
	"github.com/adityagkmit/bookstore/middlewares"
	"github.com/adityagkmit/bookstore/repositories"
	"github.com/adityagkmit/bookstore/services"

	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/mongo"
)

// SetupBookRoutes sets up book-related routes.
func SetupBookRoutes(r chi.Router, db *mongo.Database) {
	repo := repositories.NewBookRepository(db)
	service := services.NewBookService(repo)
	handler := handlers.NewBookHandler(service)

	r.Use(middlewares.AuthMiddleware) // Apply auth middleware

	r.Post("/", handler.CreateBook)
	r.Get("/", handler.GetAllBooks)
	r.Get("/{id}", handler.GetBookByID)
	r.Put("/{id}", handler.UpdateBook)
	r.Delete("/{id}", handler.DeleteBook)
}
