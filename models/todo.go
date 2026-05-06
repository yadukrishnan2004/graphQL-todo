package models

type Todo struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string
	Completed bool
	UserID    uint
	User      User `gorm:"constraint:OnDelete:CASCADE;"`
}