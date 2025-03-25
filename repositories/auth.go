package repositories

import (
	"context"
	"errors"

	"github.com/adityagkmit/bookstore/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepository struct {
	db *mongo.Collection
}

func NewAuthRepository(db *mongo.Database) *AuthRepository {
	return &AuthRepository{db: db.Collection("users")}
}

func (ar *AuthRepository) CreateUser(user *models.User) error {
	_, err := ar.db.InsertOne(context.TODO(), user)
	return err
}

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
