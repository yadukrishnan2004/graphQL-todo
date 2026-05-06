package models


type User struct {
	ID       uint   `gorm:"primaryKey"`
	Username string
	Email    string `gorm:"unique"`
	Password string
	Todos    []Todo
}