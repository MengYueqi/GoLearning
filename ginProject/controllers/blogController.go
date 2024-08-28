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
