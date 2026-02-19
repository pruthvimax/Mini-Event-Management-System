package main

import (
	"github.com/pruthvimax/user-service/database"
	"github.com/pruthvimax/user-service/handlers"
	"github.com/gin-gonic/gin"
	"github.com/pruthvimax/user-service/middleware"

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
	router.POST("/users/login", handlers.LoginUser)
	router.GET("/users/profile", middleware.AuthMiddleware(), handlers.UserProfile)


	router.Run(":8000")
}
