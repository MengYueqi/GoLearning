package main

import (
	"ginProject/dao"
	"ginProject/router"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := router.SetupRouter()
	dao.Init()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
