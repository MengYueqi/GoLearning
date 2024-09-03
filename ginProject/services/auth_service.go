package services

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// 定义一个用于存储 JWT 密钥的变量
var jwtKey = []byte("your_secret_key")

// 定义 JWT 的声明结构体
type Claims struct {
	Username string `json:"username"`
	UserId   uint   `json:"user_id"`
	jwt.StandardClaims
}

// 生成 JWT 令牌的函数
func GenerateJWT(username string, userId uint) (string, error) {
	// 设置令牌的过期时间
	expirationTime := time.Now().Add(1 * time.Hour)

	// 创建声明
	claims := &Claims{
		Username: username,
		UserId:   userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// 创建 Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 签名 Token
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求头中的Token
		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is required"})
			c.Abort()
			return
		}

		// 解析Token
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// 将用户名写入上下文
		c.Set("username", claims.Username)
		c.Next()
	}
}
