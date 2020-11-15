package model

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// DatabaseName :
	DatabaseName string = "graphql-tutorial"
	// CollectionPost :
	CollectionPost string = "Post"
)

var client *mongo.Client

// InitMongoDB :
func InitMongoDB() {
	conn := fmt.Sprintf(
		"mongodb://%s:%s/%s",
		"localhost",
		"27017",
		"graphql-tutorial",
	)

	log.Println(conn)

	c, err := mongo.NewClient(options.Client().ApplyURI(conn))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := c.Connect(ctx); err != nil {
		panic(err)
	}

	if err := c.Ping(context.TODO(), nil); err != nil {
		panic(err)
	}

	client = c
}
