package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID  `bson:"_id,omitempty" json:"id"`
	Name      string              `bson:"name" json:"name" validate:"required,min=3"`
	Email     string              `bson:"email" json:"email" validate:"required,email"`
	Password  string              `bson:"password,omitempty" json:"password" validate:"required,min=6"`
	CreatedAt time.Time           `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time           `bson:"updatedAt" json:"updatedAt"`
	DeletedAt *primitive.DateTime `bson:"deletedAt,omitempty" json:"deletedAt,omitempty"`
}

// Validate method for User model
func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
