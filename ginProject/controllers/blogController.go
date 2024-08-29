package controllers

import (
	"fmt"
	"ginProject/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BlogForm struct {
	UserId int `json:"userId"`
}

type AddBlogForm struct {
	AuthorId int    `json:"authorId"`
	Content  string `json:"content"`
}

type DeleteBlogForm struct {
	BlogId int `json:"blogId"`
}

func GetAllBlogsById(c *gin.Context) {
	var blogForm BlogForm

	// 获取并解析 JSON 数据
	if err := c.ShouldBind(&blogForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	userId := blogForm.UserId
	blogs, _ := dao.GetAllBlogsById(userId)
	fmt.Println(blogs)

	c.JSON(http.StatusOK, gin.H{
		"authorId": userId,
		"blogs":    blogs,
	})
}

func AddBlog(c *gin.Context) {
	var addBlogForm AddBlogForm

	// 获取并解析 JSON 数据
	if err := c.ShouldBind(&addBlogForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	AuthorId, Context := addBlogForm.AuthorId, addBlogForm.Content
	err := dao.AddBlog(AuthorId, Context)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
	}
}
func DeleteBlog(c *gin.Context) {
	var deleteBlogForm DeleteBlogForm

	// 获取并解析 JSON 数据
	if err := c.ShouldBind(&deleteBlogForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	blogId := deleteBlogForm.BlogId
	err := dao.DeleteBlogById(blogId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "success",
		})
	}
}
