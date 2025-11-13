package main

import (
	"github.com/gin-gonic/gin"
	"github.com/victor-lima-142/simple-financial-api/database"
)

func main() {
	router := gin.Default()
	database.InitMongoDB("", "")
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run()
}
