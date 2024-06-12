package utils

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/XDBerry29/e-commerce-go+react/domain/models"
)

func ConnectDatabase() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Automatically migrate user model into the db
	err = db.AutoMigrate(&models.User{}, &models.Product{})
	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}

	log.Println("Database connected successfully")
	return db, nil
}
