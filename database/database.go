package database

import (
	"log"
	"os"
	"service-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() {
	dsn := os.Getenv("DB_URL")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Printf("Connected to database:%v", DB.Name())

	DB.AutoMigrate(&models.Servicer{})
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Booking{})
	DB.AutoMigrate(&models.Saved{})

}
