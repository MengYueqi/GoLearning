package controllers

import (
	"fmt"
	"ginProject/dao"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "Login",
	})
}

func Login(c *gin.Context) {
	// 获取表单数据
	username := c.PostForm("username")
	password := c.PostForm("password")
	user, error := dao.GetUserByUsername(username)

	// 这里可以添加用户名和密码的验证逻辑
	if user != nil && user.Password == password {
		c.JSON(http.StatusOK, gin.H{
			"status": "Login successful",
		})
		c.Redirect(http.StatusFound, "/home")
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "Login failed",
		})
		fmt.Print(error)
		c.Redirect(http.StatusUnauthorized, "/")
	}
}
