package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/adityagkmit/bookstore/models"
	"github.com/adityagkmit/bookstore/services"

	"github.com/go-chi/chi/v5"
)

// BookHandler handles HTTP requests for books.
type BookHandler struct {
	service *services.BookService
}

// NewBookHandler initializes a new BookHandler.
func NewBookHandler(service *services.BookService) *BookHandler {
	return &BookHandler{service: service}
}

// CreateBook handles adding a new book.
func (bh *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := bh.service.CreateBook(&book); err != nil {
		http.Error(w, "Failed to create book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

// GetAllBooks handles retrieving all books.
func (bh *BookHandler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := bh.service.GetAllBooks()
	if err != nil {
		http.Error(w, "Failed to get books", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(books)
}

// GetBookByID handles retrieving a book by ID.
func (bh *BookHandler) GetBookByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	book, err := bh.service.GetBookByID(id)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(book)
}

// UpdateBook handles updating a book.
func (bh *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := bh.service.UpdateBook(id, &book); err != nil {
		http.Error(w, "Failed to update book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// DeleteBook handles deleting a book.
func (bh *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := bh.service.DeleteBook(id); err != nil {
		http.Error(w, "Failed to delete book", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
