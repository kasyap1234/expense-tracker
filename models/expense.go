package models



type Expense struct {
	ID uint `gorm:"primaryKey"`
	Title string `gorm:"not null"`
	Description string `gorm:"not null"`
	Amount float64 `gorm:"not null"`
	UserID uint `gorm:"not null"`
	User User `gorm:"foreignKey:UserID"`
}

