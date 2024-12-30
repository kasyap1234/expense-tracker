package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kasyap1234/expense-tracker/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	godotenv.Load()
	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	DB = db
	DB.AutoMigrate(&models.User{}, &models.Expense{})
}
