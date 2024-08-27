package main

import (
	"fmt"
	"ginProject/dao"
	"ginProject/router"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := router.SetupRouter()
	dao.Init()
	fmt.Print(dao.GetAllUsers())
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
