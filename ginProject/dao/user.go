package dao

import (
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
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
	// 定义 Redis 中的缓存键
	cacheKey := fmt.Sprintf("user:username:%s", username)

	// 尝试从 Redis 中获取用户数据
	userData, err := rdb.Get(ctx, cacheKey).Result()
	if err == nil {
		fmt.Println("Find in Redis!")
		// 如果在 Redis 中找到了用户数据，则反序列化并返回
		var user User
		if err := json.Unmarshal([]byte(userData), &user); err == nil {
			return &user, nil
		}
	}

	// 从数据库里获取数据
	var user User
	result := db.Where("username = ?", username).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	// 将数据添加到 Redis 并返回
	// 将查询到的用户数据序列化为 JSON 字符串
	userJSON, err := json.Marshal(user)
	// 将序列化后的用户数据缓存到 Redis 中，设置过期时间为 10 分钟
	err = rdb.Set(ctx, cacheKey, userJSON, 10*time.Minute).Err()
	if err != nil {
		fmt.Printf("Error setting cache for user %s: %v\n", username, err)
	}
	return &user, result.Error
}
