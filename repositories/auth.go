package repositories

import (
	"context"
	"errors"

	"github.com/adityagkmit/bookstore/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// AuthRepository handles user authentication-related database operations.
type AuthRepository struct {
	db *mongo.Collection
}

// NewAuthRepository initializes a new AuthRepository.
func NewAuthRepository(db *mongo.Database) *AuthRepository {
	return &AuthRepository{db: db.Collection("users")}
}

// CreateUser inserts a new user into the database.
func (ar *AuthRepository) CreateUser(user *models.User) error {
	user.SetTimestamps()
	_, err := ar.db.InsertOne(context.TODO(), user)
	return err
}

// GetUserByEmail retrieves a user by email.
func (ar *AuthRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := ar.db.FindOne(context.TODO(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return &user, nil
}
