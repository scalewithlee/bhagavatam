package database

import (
	"bhagavatam/internal/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// global variable for sharing database connections
var DB *gorm.DB

func Connect(databaseURL string) {
	var err error

	DB, err = gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	log.Print("Database connection established")
}

func Migrate() {
	err := DB.AutoMigrate(&models.Verse{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database migration complete")
}
