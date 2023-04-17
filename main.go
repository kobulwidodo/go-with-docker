package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "hello from api, update using ci cd",
		})
	})

	r.GET("/me", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "hello, its me",
		})
	})

	r.Run()
}
