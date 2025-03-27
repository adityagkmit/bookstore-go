package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID          primitive.ObjectID  `bson:"_id,omitempty" json:"id"`
	Title       string              `bson:"title" json:"title" validate:"required,min=3"`
	Author      string              `bson:"author" json:"author" validate:"required"`
	Description string              `bson:"description" json:"description"`
	Price       float64             `bson:"price" json:"price" validate:"required,gt=0"`
	CreatedAt   time.Time           `bson:"createdAt" json:"createdAt"`
	UpdatedAt   time.Time           `bson:"updatedAt" json:"updatedAt"`
	DeletedAt   *primitive.DateTime `bson:"deletedAt,omitempty" json:"deletedAt,omitempty"`
}

// Validate method for Book model
func (b *Book) Validate() error {
	validate := validator.New()
	return validate.Struct(b)
}

// SetTimestamps initializes timestamps before saving
func (b *Book) SetTimestamps() {
	currentTime := time.Now()
	if b.CreatedAt.IsZero() {
		b.CreatedAt = currentTime
	}
	b.UpdatedAt = currentTime
}
