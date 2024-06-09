package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cairomedeiros/todo-list/handler"
	"github.com/cairomedeiros/todo-list/helper"
	"github.com/cairomedeiros/todo-list/schemas"
	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	Validate *validator.Validate
}

func NewAuthHandlerImpl(validate *validator.Validate) *AuthHandler {
	return &AuthHandler{Validate: validate}
}

func (c AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	db := handler.GetDB()
	reqBody := handler.RegisterRequest{}

	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		handler.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.Validate.Struct(reqBody); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessage := fmt.Sprintf("Validation failed for field: %s", validationErrors[0].Field())
		handler.SendError(w, http.StatusBadRequest, errorMessage)
		return
	}

	var existingUser schemas.User
	result := db.Where("email = ?", reqBody.Email).First(&existingUser)
	if result.RowsAffected > 0 {
		handler.SendError(w, http.StatusBadRequest, "Email already exists")
		return
	}

	password, err := helper.EncryptPassword(reqBody.Password)

	if err != nil {
		handler.SendError(w, http.StatusBadRequest, "Failed to encrypt password")
		return
	}

	newUser := schemas.User{
		Name:     reqBody.Name,
		Email:    reqBody.Email,
		Password: password,
	}

	if err := db.Create(&newUser).Error; err != nil {
		handler.SendError(w, http.StatusBadRequest, "Failed to create user")
		return
	}

	handler.SendSuccess(w, "User registered successfully", newUser)
}

func (c AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	db := handler.GetDB()
	reqBody := handler.LoginRequest{}
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		handler.SendError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.Validate.Struct(reqBody); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessage := fmt.Sprintf("Validation failed for field: %s", validationErrors[0].Field())
		handler.SendError(w, http.StatusBadRequest, errorMessage)
		return
	}

	var existingUser schemas.User
	result := db.Where("email = ?", reqBody.Email).First(&existingUser)
	if result.RowsAffected < 1 {
		handler.SendError(w, http.StatusBadRequest, "Email not found")
		return
	}

	valid := helper.ComparePassword(reqBody.Password, existingUser.Password)

	if !valid {
		handler.SendError(w, http.StatusBadRequest, "Password invalid")
		return
	}

	token, err := helper.CreateToken(existingUser.Email)

	if err != nil {
		handler.SendError(w, http.StatusBadRequest, "JWT Error")
		return
	}

	handler.SendSuccess(w, "User logged successfully", token)
}
