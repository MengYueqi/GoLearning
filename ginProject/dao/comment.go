package dao

import (
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"time"
)

type CommentWithName struct {
	Id        int `gorm:"primary_key;AUTO_INCREMENT"`
	UserId    int
	Username  string
	BlogId    int
	Content   string
	CreatedAt time.Time
}

type Comment struct {
	Id        int `gorm:"primary_key;AUTO_INCREMENT"`
	UserId    int
	BlogId    int
	Content   string
	CreatedAt time.Time
}

// GetAllCommentOnOneBlog 根据 Id 获取博客的所有评论
func GetAllCommentOnOneBlog(blogId int) (*[]CommentWithName, error) {
	// 定义 Redis 中的缓存键
	cacheKey := fmt.Sprintf("blog:blogId:%d", blogId)

	// 尝试从 Redis 中获取用户数据
	commentsData, err := rdb.Get(ctx, cacheKey).Result()
	var comments *[]CommentWithName
	if err == nil {
		fmt.Println("Find in Redis!")
		// 如果在 Redis 中找到了用户数据，则反序列化并返回
		if err := json.Unmarshal([]byte(commentsData), &comments); err == nil {
			return comments, nil
		}
	}

	result := db.Table("comments").Select("comments.id, comments.user_id, users.username, comments.blog_id, comments.content, comments.created_at").
		Joins("JOIN users on users.user_id = comments.user_id").Where("blog_id = ?", blogId).Find(&comments)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	// 将数据添加到 Redis 并返回
	// 将查询到的用户数据序列化为 JSON 字符串
	commentsJSON, err := json.Marshal(comments)
	// 将序列化后的用户数据缓存到 Redis 中，设置过期时间为 10 分钟
	err = rdb.Set(ctx, cacheKey, commentsJSON, 10*time.Minute).Err()
	if err != nil {
		fmt.Printf("Error setting cache for comments for BLOG %d: %v\n", blogId, err)
	}
	return comments, result.Error
}

// AddCommentById 增加一个评论
func AddCommentById(blogId int, userId int, content string) error {
	comment := Comment{
		UserId:  userId,
		BlogId:  blogId,
		Content: content,
	}
	result := db.Create(&comment)

	// 删除缓存中对应 Blog 的评论数据
	cacheKey := fmt.Sprintf("blog:blogId:%d", blogId)
	_, err := rdb.Del(ctx, cacheKey).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	return result.Error
}
