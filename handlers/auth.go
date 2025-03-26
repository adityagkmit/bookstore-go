package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/adityagkmit/bookstore/models"
	"github.com/adityagkmit/bookstore/services"
	"github.com/adityagkmit/bookstore/utils"
)

type AuthHandler struct {
	service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func (ah *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid input")
		return
	}

	token, err := ah.service.Register(&user)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendSuccessResponse(w, map[string]string{"token": token}, "User registered successfully")
}

func (ah *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var credentials models.User
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid input")
		return
	}

	token, err := ah.service.Login(credentials.Email, credentials.Password)
	if err != nil {
		utils.SendErrorResponse(w, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	utils.SendSuccessResponse(w, map[string]string{"token": token}, "Login successful")
}
