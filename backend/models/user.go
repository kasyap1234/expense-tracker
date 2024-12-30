package models 

type User struct {
	ID uint `gorm:"primaryKey"`
	Username string `gorm:"not null;unique"`
	Email string `gorm:"not null;unique"`
    
}

