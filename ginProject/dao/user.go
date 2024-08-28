package dao

import (
	"errors"
	"gorm.io/gorm"
)

// User 模型
type User struct {
	User_id  uint `gorm:"primaryKey"`
	Username string
	Password string
	Email    string
}

func GetAllUsers() ([]User, error) {
	var users []User
	result := db.Find(&users)
	return users, result.Error
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	result := db.Where("username = ?", username).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, result.Error
}
