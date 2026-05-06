package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	var database *gorm.DB
	var err error

	for i := 0; i < 10; i++ {

		database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err == nil {
			fmt.Println("Database connected")
			DB = database
			return
		}

		fmt.Println("Waiting for database...")
		time.Sleep(2 * time.Second)
	}

	log.Fatal("Failed to connect DB:", err)
}