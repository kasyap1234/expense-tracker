package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/kasyap1234/expense-tracker/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	godotenv.Load()
	dsn := os.Getenv("DB_URL")
	
	maxRetries := 5
	for i := 0; i < maxRetries; i++ {
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			DB = db
			DB.AutoMigrate(&models.User{}, &models.Expense{})
			log.Printf("Successfully connected to database")
			return
		}
		log.Printf("Failed to connect to database, attempt %d/%d: %v", i+1, maxRetries, err)
		time.Sleep(5 * time.Second)
	}
	log.Fatal("Failed to connect to database after multiple attempts")
}
