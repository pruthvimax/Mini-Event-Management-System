package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func connectDB() {
	clientOptions := options.Client().ApplyURI("mongodb+srv://pruthvialalliprivate_db_user:7rexYduB9cbIL0jP@eventmanagement.oducztq.mongodb.net/?appName=EventManagement")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database("eventdb").Collection("users")
	fmt.Println("Connected to MongoDB Atlas 🚀")
}

func main() {
	connectDB()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "User Service Running",
		})
	})

	router.Run(":8000")
}
