package main

import (
	"github.com/pruthvimax/user-service/database"
	"github.com/pruthvimax/user-service/handlers"
	"github.com/gin-gonic/gin"
)


func main() {
	database.ConnectDB()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "User Service Running",
		})
	})

	router.POST("/users/register", handlers.RegisterUser)

	router.Run(":8000")
}
