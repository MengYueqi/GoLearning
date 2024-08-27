package dao

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
