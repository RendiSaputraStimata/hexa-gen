package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type {{.Domain}} struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
}