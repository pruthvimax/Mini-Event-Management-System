package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pruthvimax/event-service/database"
	"github.com/pruthvimax/event-service/handlers"
	"github.com/pruthvimax/event-service/middleware"
)

func main() {

	database.ConnectDB()

	router := gin.Default()

	router.POST("/events/create", middleware.AuthMiddleware(), handlers.CreateEvent)
	router.GET("/events", handlers.GetEvents)

	router.Run(":8001")
}