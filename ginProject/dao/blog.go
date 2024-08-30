package dao

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type BlogsWithName struct {
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

// AddBlog 增加一个 Blog
func AddBlog(AuthorId int, Content string) error {
	blog := Blogs{
		AuthorId: AuthorId,
		Content:  Content,
	}
	result := db.Create(&blog)
	return result.Error
}

// GetAllBlogsById 根据 Id 获取所有的 Blogs
func GetAllBlogsById(AuthorId int) ([]*BlogsWithName, error) {
	var blogs []*BlogsWithName
	//result := db.Where("author_id = ?", AuthorId).Find(&blogs)
	result := db.Table("blogs").Select("blogs.id, blogs.created_at, blogs.author_id, blogs.content, users.username").
		Joins("JOIN users ON users.user_id = blogs.author_id").Where("blogs.author_id = ?", AuthorId).Find(&blogs)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return blogs, result.Error
}

func GetAllBlogs() ([]*BlogsWithName, error) {
	var blogs []*BlogsWithName
	result := db.Table("blogs").Select("blogs.id, blogs.created_at, blogs.author_id, blogs.content, users.username").
		Joins("JOIN users ON users.user_id = blogs.author_id").Find(&blogs)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return blogs, result.Error
}

// DeleteBlogById 根据 Id 删除相关的 Blog
func DeleteBlogById(BlogId int) error {
	result := db.Where("id = ?", BlogId).Delete(&Blogs{})
	return result.Error
}

// ModifyBlogById 根据 Id 更改相关的 Blog
func ModifyBlogById(BlogId int, Content string) error {
	result := db.Model(&Blogs{}).Where("id = ?", BlogId).Update("content", Content)
	return result.Error
}
