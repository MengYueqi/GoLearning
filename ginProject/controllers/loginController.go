package controllers

import (
	"fmt"
	"ginProject/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Login",
	})
}

func Login(c *gin.Context) {
	var loginReq LoginRequest

	// 获取并解析 JSON 数据
	if err := c.ShouldBind(&loginReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// 获取用户名和密码
	username := loginReq.Username
	password := loginReq.Password
	fmt.Println(username, password)

	user, error := dao.GetUserByUsername(username)

	// 这里可以添加用户名和密码的验证逻辑
	if user != nil && user.Password == password {
		c.JSON(http.StatusOK, gin.H{
			"status": "Login successful",
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "Login failed",
		})
		fmt.Print(error)
	}
}
