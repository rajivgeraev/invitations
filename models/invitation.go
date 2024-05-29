package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Invitation struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Code      string             `bson:"code"`
	MaxUses   int                `bson:"max_uses"`
	UsedCount int                `bson:"used_count"`
	Emails    []string           `bson:"emails"`
}
