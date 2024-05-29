package services

import (
	"context"
	"errors"
	"time"

	"github.com/rajivgeraev/config"
	"github.com/rajivgeraev/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func RegisterUser(code, email string) error {
	collection := config.MongoDB.Collection("invitations")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"code": code}
	update := bson.M{
		"$addToSet": bson.M{"emails": email},
		"$inc":      bson.M{"used_count": 1},
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After).SetUpsert(false)
	var updatedInvitation models.Invitation

	err := collection.FindOneAndUpdate(ctx, filter, update, opts).Decode(&updatedInvitation)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New("invitation code not found")
		}
		return err
	}

	if updatedInvitation.UsedCount > updatedInvitation.MaxUses {
		return errors.New("invitation code has been used maximum times")
	}

	return nil
}
