package controllers

import (
	"fmt"
	"ginProject/dao"
	"ginProject/services"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

var jwtKey = []byte("your_secret_key") // 替换为自己的密钥

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Claims struct {
	Username string `json:"username"`
	UserId   uint   `json:"user_id"`
	jwt.StandardClaims
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

	user, _ := dao.GetUserByUsername(username)

	// 使用 JWT 鉴权
	tokenString, err := services.GenerateJWT(user.Username, user.User_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create token"})
		return
	}

	// 返回Token
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}
