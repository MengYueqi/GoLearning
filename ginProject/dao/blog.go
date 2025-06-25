package dao

import (
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"strconv"
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

type BlogLike struct {
	BlogId int
	UserId int
}

// AddBlog 增加一个 Blog
func AddBlog(AuthorId int, Content string) error {
	blog := Blogs{
		AuthorId: AuthorId,
		Content:  Content,
	}
	result := db.Create(&blog)
	// 定义 Redis 里的缓存键
	cacheKey := fmt.Sprintf("AllBlogs:")
	_, err := rdb.Del(ctx, cacheKey).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
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

// LikeBlog 根据 Id 对 Blog 进行点赞
func LikeBlog(blogId int, userId int) error {
	blogLike := BlogLike{
		BlogId: blogId,
		UserId: userId,
	}
	err := db.Create(&blogLike).Error
	if err != nil {
		return err
	} else {
		return nil
	}
}

// GetBlogLikesById 根据 Id 获取 Blog 的点赞数
func GetBlogLikesById(blogId int) (int64, error) {
	var count int64

	key := fmt.Sprintf("blog:likes:%d", blogId)
	// 从 Redis 获取缓存
	likesStr, err := rdb.Get(ctx, key).Result()
	if err == nil {
		// 成功命中缓存
		count, convErr := strconv.ParseInt(likesStr, 10, 64)
		fmt.Println(key + " Like Num Find in Redis!")
		if convErr == nil {
			return count, nil
		}
	}

	// 解析失败就继续查数据库
	result := db.Model(&BlogLike{}).Where("blog_id = ?", blogId).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}

	// 将结果写入 Redis，设置过期时间 10 分钟
	rdb.Set(ctx, key, strconv.FormatInt(count, 10), 10*time.Minute)

	return count, nil
}

// GetAllBlogs 获取所有博客
func GetAllBlogs() ([]*BlogsWithName, error) {
	// 定义 Redis 里的缓存键
	cacheKey := fmt.Sprintf("AllBlogs:")

	// 在 Redis 中获取相关数据
	var blogs []*BlogsWithName
	allBlogs, err := rdb.Get(ctx, cacheKey).Result()
	if err == nil {
		fmt.Println("Find in Redis!")
		// 如果在 Redis 里获取了数据，则对数据进行反序列化
		if err := json.Unmarshal([]byte(allBlogs), &blogs); err != nil {
			return blogs, err
		}
	}

	result := db.Table("blogs").Select("blogs.id, blogs.created_at, blogs.author_id, blogs.content, users.username").
		Joins("JOIN users ON users.user_id = blogs.author_id").Find(&blogs)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	// 将数据添加到 Redis 中并返回
	// 将数据序列化为 JSON 字符串
	blogsJSON, err := json.Marshal(blogs)
	// 将序列化后的用户数据缓存到 Redis 中，设置过期时间为 10 分钟
	err = rdb.Set(ctx, cacheKey, blogsJSON, 10*time.Minute).Err()
	if err != nil {
		fmt.Printf("Error setting cache for all blogs: %v\n", err)
	}
	return blogs, result.Error
}

// DeleteBlogById 根据 Id 删除相关的 Blog
func DeleteBlogById(BlogId int) error {
	result := db.Where("id = ?", BlogId).Delete(&Blogs{})
	// 定义 Redis 里的缓存键
	cacheKey := fmt.Sprintf("AllBlogs:")
	_, err := rdb.Del(ctx, cacheKey).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	return result.Error
}

// ModifyBlogById 根据 Id 更改相关的 Blog
func ModifyBlogById(BlogId int, Content string) error {
	result := db.Model(&Blogs{}).Where("id = ?", BlogId).Update("content", Content)
	// 定义 Redis 里的缓存键
	cacheKey := fmt.Sprintf("AllBlogs:")
	_, err := rdb.Del(ctx, cacheKey).Result()
	if err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	}
	return result.Error
}
