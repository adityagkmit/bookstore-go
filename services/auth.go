package services

import (
	"errors"

	"github.com/adityagkmit/bookstore/models"
	"github.com/adityagkmit/bookstore/repositories"
	"github.com/adityagkmit/bookstore/utils"
)

type AuthService struct {
	repo *repositories.AuthRepository
}

func NewAuthService(repo *repositories.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (as *AuthService) Register(user *models.User) (string, error) {
	// Check if user already exists
	existingUser, err := as.repo.GetUserByEmail(user.Email)

	if err == nil && existingUser != nil {
		return "", errors.New("user with this email already exists")
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return "", err
	}
	user.Password = hashedPassword

	// Create the user
	err = as.repo.CreateUser(user)
	if err != nil {
		return "", err
	}

	// Generate and return JWT token
	return utils.GenerateJWT(user.Email)
}

func (as *AuthService) Login(email, password string) (string, error) {
	user, err := as.repo.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	if !utils.VerifyPassword(user.Password, password) {
		return "", errors.New("invalid credentials")
	}

	return utils.GenerateJWT(user.Email)
}
