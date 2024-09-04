package router

import (
	"ginProject/controllers"
	"ginProject/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // 允许所有域（生产环境中应根据需要进行限制）
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type, X-CSRF-Token")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}

func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	r.Use(corsMiddleware())
	//r.LoadHTMLGlob("templates/*")

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Login 对应的路由
	r.GET("/", controllers.LoginPage)
	// HomePage 对应的路由
	r.GET("/home", controllers.HomePage)
	// 登录检查
	r.POST("/login", controllers.Login)

	// 博客操作
	r.POST("/GetAllBlogsById", services.JWTAuthMiddleware(), controllers.GetAllBlogsById)
	r.POST("/addBlog", controllers.AddBlog)
	r.POST("/deleteBlog", controllers.DeleteBlog)
	r.POST("/getAllBlogs", services.JWTAuthMiddleware(), controllers.GetAllBlogs)
	r.POST("/modifyBlogById", controllers.ModifyBlogById)

	// 评论操作
	r.POST("/getAllCommentsById", controllers.GetAllCommentsById)
	r.POST("/addCommentById", controllers.AddCommentById)

	return r
}
