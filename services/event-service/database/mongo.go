package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var EventCollection *mongo.Collection

func ConnectDB() {

	clientOptions := options.Client().ApplyURI("mongodb+srv://pruthvialalliprivate_db_user:7rexYduB9cbIL0jP@eventmanagement.oducztq.mongodb.net/?appName=EventManagement")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	EventCollection = client.Database("eventdb").Collection("events")
	log.Println("Event Service Connected to MongoDB")
}