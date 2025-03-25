package services

import (
	"errors"
	"os"
	"time"

	"github.com/adityagkmit/bookstore/models"
	"github.com/adityagkmit/bookstore/repositories"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo *repositories.AuthRepository
}

func NewAuthService(repo *repositories.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

// HashPassword hashes a given password
func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed), err
}

// VerifyPassword checks if the provided password is correct
func VerifyPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// GenerateJWT generates a JWT token
func GenerateJWT(email string) (string, error) {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		return "", errors.New("JWT_SECRET is not set")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func (as *AuthService) Register(user *models.User) (string, error) {
	// Check if user already exists
	existingUser, err := as.repo.GetUserByEmail(user.Email)

	if err == nil && existingUser != nil {
		return "", errors.New("user with this email already exists")
	}

	// Hash the password
	hashedPassword, err := HashPassword(user.Password)
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
	return GenerateJWT(user.Email)
}

func (as *AuthService) Login(email, password string) (string, error) {
	user, err := as.repo.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	if !VerifyPassword(user.Password, password) {
		return "", errors.New("invalid credentials")
	}

	return GenerateJWT(user.Email)
}
