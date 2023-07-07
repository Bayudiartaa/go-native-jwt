package controllers

import (
	"encoding/json"
	"errors"
	"go-native-jwt/configs"
	"go-native-jwt/helpers"
	"go-native-jwt/models"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var register models.Register

	if err := json.NewDecoder(r.Body).Decode(&register); err != nil {
        helpers.Response(w, 500, err.Error(), nil)
		return 
	}

	defer r.Body.Close()

	if register.Name == "" {
		helpers.Response(w, 400, "Name cannot be empty", nil)
		return
	}

	if register.Password != register.PasswordConfirmation {
		helpers.Response(w, 400, "Password Not Match", nil)
		return 
	}

	// Check if email is already registered
	if err := isEmailRegistered(register.Email); err != nil {
		helpers.Response(w, 400, err.Error(), nil)
		return
	}

	passwordHash, err := helpers.HashPassword(register.Password)
	if err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	user := models.User{
		Name:     register.Name,
		Email:    register.Email,
		Password: passwordHash,
	}

	if err := configs.DB.Create(&user).Error; err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 201, "Register Successfully", nil)
	
}

func Login(w http.ResponseWriter, r *http.Request) {
	var login models.Login

	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	var user models.User
	if err := configs.DB.First(&user, "email = ?", login.Email).Error; err != nil {
		helpers.Response(w, 404, "Wrong email or password", nil)
		return
	}

	if err := helpers.VerifyPassword(user.Password, login.Password); err != nil {
		helpers.Response(w, 404, "Wrong email or password", nil)
		return
	}

	token, err := helpers.CreateToken(&user)
	if err != nil {
		helpers.Response(w, 500, err.Error(), nil)
		return
	}

	helpers.Response(w, 200, "Successfully Login", token)
}

func isEmailRegistered(email string) error {
	var user models.User
	result := configs.DB.Where("email = ?", email).First(&user)
	if result.RowsAffected > 0 {
		return errors.New("email is already registered")
	}
	return nil
}