package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/adityagkmit/bookstore/models"
	"github.com/adityagkmit/bookstore/services"
	"github.com/adityagkmit/bookstore/utils"
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
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := book.Validate(); err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := bh.service.CreateBook(&book); err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to create book")
		return
	}

	utils.SendSuccessResponse(w, book, "Book created successfully")
}

// GetAllBooks handles retrieving all books.
func (bh *BookHandler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := bh.service.GetAllBooks()
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to get books")
		return
	}

	utils.SendSuccessResponse(w, books, "Books retrieved successfully")
}

// GetBookByID handles retrieving a book by ID.
func (bh *BookHandler) GetBookByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	book, err := bh.service.GetBookByID(id)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusNotFound, "Book not found")
		return
	}

	utils.SendSuccessResponse(w, book, "Book retrieved successfully")
}

// UpdateBook handles updating a book.
func (bh *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var book models.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	if err := book.Validate(); err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := bh.service.UpdateBook(id, &book); err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to update book")
		return
	}
	utils.SendSuccessResponse(w, book, "Book updated successfully")
}

// DeleteBook handles deleting a book.
func (bh *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if err := bh.service.DeleteBook(id); err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to delete book")
		return
	}

	utils.SendSuccessResponse(w, nil, "Book deleted successfully")
}
