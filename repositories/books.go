package repositories

import (
	"context"
	"fmt"

	"github.com/adityagkmit/bookstore/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// BookRepository handles database operations for books.
type BookRepository struct {
	collection *mongo.Collection
}

// NewBookRepository initializes a new BookRepository.
func NewBookRepository(db *mongo.Database) *BookRepository {
	return &BookRepository{
		collection: db.Collection("books"),
	}
}

// CreateBook inserts a new book into the database.
func (r *BookRepository) CreateBook(book *models.Book) error {
	book.ID = primitive.NewObjectID()
	_, err := r.collection.InsertOne(context.TODO(), book)
	return err
}

// GetAllBooks retrieves all books from the database.
func (r *BookRepository) GetAllBooks() ([]models.Book, error) {
	var books []models.Book
	cursor, err := r.collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var book models.Book
		if err := cursor.Decode(&book); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

// GetBookByID retrieves a book by its ID.
func (r *BookRepository) GetBookByID(id string) (*models.Book, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	var book models.Book
	err = r.collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&book)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

// UpdateBook updates an existing book.
func (r *BookRepository) UpdateBook(id string, book *models.Book) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid book ID format")
	}

	// Perform the update
	result, err := r.collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": objectID},
		bson.M{"$set": book},
	)

	// Check if any document was updated
	if result.MatchedCount == 0 {
		return fmt.Errorf("book with ID %s not found", id)
	}

	if err != nil {
		return fmt.Errorf("failed to update book: %w", err)
	}

	// Set the correct ID in the book response
	book.ID = objectID

	return nil
}

// DeleteBook removes a book from the database.
func (r *BookRepository) DeleteBook(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.collection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
	return err
}
