package config 

import (
	"github.com/kasyap1234/expense-tracker/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB 

func InitDB(DB_URL string){
	dsn :=os.Getenv(DB_URL); 
	var err error 
	DB,err = gorm.Open(postgres.Open(dsn),&gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database")
	}
	DB.AutoMigrate(&models.User{},&models.Expense{})
	

}

