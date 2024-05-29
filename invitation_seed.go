package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/rajivgeraev/invitations/config"
	"github.com/rajivgeraev/invitations/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func main() {
	// Подключение к базе данных
	mongoURL := os.Getenv("MONGO_URL")
	if mongoURL == "" {
		log.Fatal("MONGO_URL is not set in environment variables")
	}

	config.ConnectDB()

	collection := config.MongoDB.Collection("invitations")

	invitations := []models.Invitation{
		{ID: primitive.NewObjectID(), Code: "twitter-reg1", MaxUses: 100, UsedCount: 0, Emails: []string{}},
		{ID: primitive.NewObjectID(), Code: "telegram-test", MaxUses: 100, UsedCount: 0, Emails: []string{}},
		{ID: primitive.NewObjectID(), Code: "instagram-hello", MaxUses: 100, UsedCount: 0, Emails: []string{}},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for _, invitation := range invitations {
		_, err := collection.InsertOne(ctx, invitation)
		if err != nil {
			log.Fatal("Error inserting invitation: ", err)
		}
	}

	log.Println("Data seeded successfully!")
}
