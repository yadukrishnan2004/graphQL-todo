package service

import (
	"os"
	"testing"
	"todo/db"
	"todo/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


func init() {
	os.Setenv("JWT_SECRET", "supersecret")
}

func setupTestDB() {

	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect test db")
	}

	db.DB = database

	db.DB.AutoMigrate(&models.User{})
}

func TestMain(m *testing.M) {

	setupTestDB()

	code := m.Run()

	os.Remove("test.db")

	os.Exit(code)
}


func TestRegister(t *testing.T) {

	token, err := Register(
		"testuser",
		"test@mail.com",
		"password123",
	)

	if err != nil {
		t.Errorf("register failed: %v", err)
	}

	if token == "" {
		t.Errorf("token is empty")
	}
}


func TestLogin(t *testing.T) {

	_, _ = Register(
		"loginuser",
		"login@mail.com",
		"password123",
	)

	token, err := Login(
		"login@mail.com",
		"password123",
	)

	if err != nil {
		t.Errorf("login failed: %v", err)
	}

	if token == "" {
		t.Errorf("token is empty")
	}
}


func TestLoginWrongPassword(t *testing.T) {

	_, _ = Register("wrongpassuser","wrong@mail.com","password123")

	_, err := Login("wrong@mail.com","wrongpassword")

	if err == nil {
		t.Errorf("expected login failure")
	}
}