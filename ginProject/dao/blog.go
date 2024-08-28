package dao

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Blogs struct {
	Id        int `gorm:"primary_key;AUTO_INCREMENT"`
	CreatedAt time.Time
	AuthorId  int
	Content   string
}

func GetAllBlogsById(AuthorId int) ([]*Blogs, error) {
	var blogs []*Blogs
	fmt.Println(AuthorId)
	result := db.Where("author_id = ?", AuthorId).Find(&blogs)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return blogs, result.Error
}
