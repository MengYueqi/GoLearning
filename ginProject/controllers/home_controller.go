package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomePage(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", gin.H{
		"title": "Home Page",
	})
}
