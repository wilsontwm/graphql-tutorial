package model

import (
	"context"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Post :
type Post struct {
	ID        primitive.ObjectID `bson:"_id"`
	Title     string             `bson:"title"`
	Content   string             `bson:"content"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}

// ListPost :
func ListPost(ctx context.Context, limit int, cursor string) ([]*Post, string, error) {
	query := bson.M{}
	posts := make([]*Post, 0)
	currentSkip := int64(0)

	if limit <= 0 {
		return nil, "", fmt.Errorf("Limit must be more than zero")
	}

	if cursor != "" {
		data, err := hex.DecodeString(cursor)
		if err != nil {
			return nil, "", err
		}

		currentSkip, err = strconv.ParseInt(string(data), 10, 64)
		if err != nil {
			return nil, "", err
		}
	}

	nextCursor, err := client.Database(DatabaseName).Collection(CollectionPost).Find(
		ctx,
		query,
		options.Find().SetLimit(int64(limit+1)).SetSkip(currentSkip),
	)

	defer nextCursor.Close(ctx)
	if err != nil {
		return nil, "", err
	}

	for nextCursor.Next(ctx) {
		post := new(Post)
		if err := nextCursor.Decode(post); err != nil {
			return nil, "", err
		}
		posts = append(posts, post)
	}

	if err := nextCursor.Err(); err != nil {
		return nil, "", err
	}

	if len(posts) > int(limit) {
		return posts[:len(posts)-1], hex.EncodeToString([]byte(fmt.Sprintf("%d", currentSkip+int64(limit)))), nil
	}
	return posts, "", nil
}

// GetPost :
func GetPost(ctx context.Context, id string) (*Post, error) {
	p := new(Post)

	hexID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	if err := client.Database(DatabaseName).Collection(CollectionPost).FindOne(
		ctx,
		bson.M{
			"_id": hexID,
		},
	).Decode(p); err != nil {
		return nil, err
	}

	return p, nil
}

// CreatePost :
func CreatePost(ctx context.Context, post *Post) error {
	collection := client.Database(DatabaseName).Collection(CollectionPost)

	_, err := collection.InsertOne(context.TODO(), post)

	return err
}

// UpdatePost :
func UpdatePost(ctx context.Context, post *Post) error {
	_, err := client.Database(DatabaseName).Collection(CollectionPost).UpdateOne(
		ctx,
		bson.M{"_id": post.ID}, // Filter
		bson.M{"$set": post},   // Update
	)

	return err
}

// DeletePost :
func DeletePost(ctx context.Context, post *Post) error {
	_, err := client.Database(DatabaseName).Collection(CollectionPost).DeleteOne(
		ctx,
		bson.M{"_id": post.ID}, // Filter
	)

	return err
}
