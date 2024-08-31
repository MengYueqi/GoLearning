package controllers

import (
	"ginProject/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

type GetAllCommentsForm struct {
	BlogId int `json:"blog_id"`
}

func GetAllCommentsById(c *gin.Context) {
	var getAllCommentsForm GetAllCommentsForm
	// 获取并解析 JSON 数据
	if err := c.ShouldBind(&getAllCommentsForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}
	comments, err := dao.GetAllCommentOnOneBlog(getAllCommentsForm.BlogId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":   "success",
			"comments": comments,
		})
	}
}
