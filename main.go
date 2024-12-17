package main

import (
	"data_crawler/module/swift_code"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	swift_code.SetupService(v1)
	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "pong"})
	})
	router.Run(":8080")
}
