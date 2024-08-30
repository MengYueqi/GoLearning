package router

import (
	"ginProject/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")

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
	r.LoadHTMLGlob("templates/*")

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
	r.POST("/GetAllBlogsById", controllers.GetAllBlogsById)
	r.POST("/addBlog", controllers.AddBlog)
	r.POST("/deleteBlog", controllers.DeleteBlog)
	r.POST("/getAllBlogs", controllers.GetAllBlogs)
	r.POST("/modifyBlogById", controllers.ModifyBlogById)

	return r
}
