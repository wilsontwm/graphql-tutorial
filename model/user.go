package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// User :
type User struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `bson:"name"`
}
