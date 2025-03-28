package repositories

import (
	"context"
	"fmt"
	"time"

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

// CreateBook inserts a new book into the database with timestamps.
func (r *BookRepository) CreateBook(book *models.Book) error {
	book.SetTimestamps()

	result, err := r.collection.InsertOne(context.TODO(), book)
	if err != nil {
		return err
	}

	// Retrieve the inserted ID
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		book.ID = oid
	}

	return nil
}

// GetAllBooks retrieves all books from the database.
func (r *BookRepository) GetAllBooks() ([]models.Book, error) {
	var books []models.Book
	cursor, err := r.collection.Find(context.TODO(), bson.M{"deletedAt": nil}) // Exclude soft-deleted books
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
	err = r.collection.FindOne(context.TODO(), bson.M{"_id": objectID, "deletedAt": nil}).Decode(&book)
	if err != nil {
		return nil, err
	}

	return &book, nil
}

// UpdateBook updates an existing book and sets the updatedAt timestamp.
func (r *BookRepository) UpdateBook(id string, book *models.Book) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("invalid book ID format")
	}

	book.SetTimestamps()

	// Perform the update, excluding CreatedAt from modification
	updateFields := bson.M{
		"title":       book.Title,
		"author":      book.Author,
		"description": book.Description,
		"price":       book.Price,
		"updatedAt":   book.UpdatedAt,
	}

	result, err := r.collection.UpdateOne(
		context.TODO(),
		bson.M{"_id": objectID, "deletedAt": nil}, // Ensure not updating deleted books
		bson.M{"$set": updateFields},
	)

	// Check if any document was updated
	if result.MatchedCount == 0 {
		return fmt.Errorf("book with ID %s not found", id)
	}

	if err != nil {
		return err
	}

	book.ID = objectID

	return nil
}

// DeleteBook marks a book as deleted instead of removing it.
func (r *BookRepository) DeleteBook(id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	// Set DeletedAt to current time instead of deleting the document
	update := bson.M{
		"$set": bson.M{"deletedAt": primitive.NewDateTimeFromTime(time.Now())},
	}

	_, err = r.collection.UpdateOne(context.TODO(), bson.M{"_id": objectID}, update)
	return err
}
