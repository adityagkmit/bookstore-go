package services

import (
	"github.com/adityagkmit/bookstore/models"
	"github.com/adityagkmit/bookstore/repositories"
)

// BookService handles business logic for books.
type BookService struct {
	repo *repositories.BookRepository
}

// NewBookService initializes a new BookService.
func NewBookService(repo *repositories.BookRepository) *BookService {
	return &BookService{repo: repo}
}

// CreateBook calls the repository to create a book.
func (bs *BookService) CreateBook(book *models.Book) error {
	return bs.repo.CreateBook(book)
}

// GetAllBooks retrieves all books.
func (bs *BookService) GetAllBooks() ([]models.Book, error) {
	return bs.repo.GetAllBooks()
}

// GetBookByID retrieves a book by its ID.
func (bs *BookService) GetBookByID(id string) (*models.Book, error) {
	return bs.repo.GetBookByID(id)
}

// UpdateBook updates an existing book.
func (bs *BookService) UpdateBook(id string, book *models.Book) error {
	return bs.repo.UpdateBook(id, book)
}

// DeleteBook removes a book from the database.
func (bs *BookService) DeleteBook(id string) error {
	return bs.repo.DeleteBook(id)
}
