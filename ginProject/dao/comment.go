package dao

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type CommentWithName struct {
	Id        int `gorm:"primary_key"`
	UserId    int
	Username  string
	BlogId    int
	Content   string
	CreatedAt time.Time
}

func GetAllCommentOnOneBlog(blogId int) (*[]CommentWithName, error) {
	var comments *[]CommentWithName
	result := db.Table("comments").Select("comments.id, comments.user_id, users.username, comments.blog_id, comments.content, comments.created_at").
		Joins("JOIN users on users.user_id = comments.user_id").Where("blog_id = ?", blogId).Find(&comments)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return comments, result.Error
}
