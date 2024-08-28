package router

import (
	"ginProject/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Login 对应的路由
	r.GET("/", controllers.LoginPage)
	// HomePage 对应的路由
	r.GET("/home", controllers.HomePage)
	r.POST("/login", controllers.Login)

	return r
}
