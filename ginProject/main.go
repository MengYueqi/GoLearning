package main

import (
	"ginProject/dao"
	"ginProject/router"
)

func main() {
	r := router.SetupRouter()
	dao.InitMySQL()
	// Listen and Server in 0.0.0.0:8081
	r.Run(":8081")
}
