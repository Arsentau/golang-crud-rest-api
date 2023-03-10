package models

type Movie struct {
	// gorm.Model
	ID          string `gorm:"primaryKey not null; unique_index"`
	Title       string `gorm:"not null"`
	Bought      bool   `gorm:"default:false"`
	Description string
}
