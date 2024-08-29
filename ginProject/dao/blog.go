package dao

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type BlogsWithId struct {
	Id        int `gorm:"primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time
	AuthorId  int
	Username  string
	Content   string
}

type Blogs struct {
	Id        int `gorm:"primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time
	AuthorId  int
	Content   string
}

func AddBlog(AuthorId int, Content string) error {
	blog := Blogs{
		AuthorId: AuthorId,
		Content:  Content,
	}
	result := db.Create(&blog)
	return result.Error
}

func GetAllBlogsById(AuthorId int) ([]*BlogsWithId, error) {
	var blogs []*BlogsWithId
	fmt.Println(AuthorId)
	//result := db.Where("author_id = ?", AuthorId).Find(&blogs)
	result := db.Table("blogs").Select("blogs.id, blogs.created_at, blogs.author_id, blogs.content, users.username").
		Joins("JOIN users ON users.user_id = blogs.author_id").Where("blogs.author_id = ?", AuthorId).Find(&blogs)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	fmt.Println(blogs[0].Username)
	return blogs, result.Error
}

func DeleteBlogById(BlogId int) error {
	result := db.Where("id = ?", BlogId).Delete(&Blogs{})
	return result.Error
}
