package models

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title       string             `bson:"title" json:"title" validate:"required,min=3"`
	Author      string             `bson:"author" json:"author" validate:"required"`
	Description string             `bson:"description" json:"description"`
	Price       float64            `bson:"price" json:"price" validate:"required,gt=0"`
}

// Validate method for Book model
func (b *Book) Validate() error {
	validate := validator.New()
	return validate.Struct(b)
}
