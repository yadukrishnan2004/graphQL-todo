package services

import (
	"errors"
	"todo/db"
	"todo/models"
	"todo/utils"
)

func Register(username, email, password string) (string, error) {
	hashed, err := utils.HashPassword(password)
	if err != nil {
		return "", err
	}

	user := models.User{
		Username: username,
		Email:    email,
		Password: hashed,
	}

	if err := db.DB.Create(&user).Error; err != nil {
		return "", err
	}

	token, err := utils.GenerateToken(user.ID)
	return token, err
}

func Login(email, password string) (string, error) {
	var user models.User

	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return "", errors.New("invalid credentials")
	}

	if !utils.CheckPassword(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	return utils.GenerateToken(user.ID)
}