package dao

import (
	"errors"
	"gorm.io/gorm"
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
	var comments *[]CommentWithName
	result := db.Table("comments").Select("comments.id, comments.user_id, users.username, comments.blog_id, comments.content, comments.created_at").
		Joins("JOIN users on users.user_id = comments.user_id").Where("blog_id = ?", blogId).Find(&comments)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
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
	return result.Error
}
