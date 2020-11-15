package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/wilsontwm/graphql-tutorial/graph/generated"
	"github.com/wilsontwm/graphql-tutorial/graph/model"
	bmodel "github.com/wilsontwm/graphql-tutorial/model"
	"github.com/wilsontwm/graphql-tutorial/model/transformer"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *mutationResolver) CreatePost(ctx context.Context, input model.CreatePostRequest) (*model.Post, error) {
	post := new(bmodel.Post)
	post.ID = primitive.NewObjectID()
	post.Title = input.Title
	post.Content = input.Content
	post.CreatedAt = time.Now().UTC()
	post.UpdatedAt = time.Now().UTC()

	if err := bmodel.CreatePost(ctx, post); err != nil {
		return nil, err
	}

	return transformer.ToPost(post), nil
}

func (r *mutationResolver) UpdatePost(ctx context.Context, input model.UpdatePostRequest) (*model.Post, error) {
	p, err := bmodel.GetPost(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	if p == nil {
		return nil, fmt.Errorf("Post not found")
	}

	p.Title = input.Title
	p.Content = input.Content
	p.UpdatedAt = time.Now().UTC()

	if err := bmodel.UpdatePost(ctx, p); err != nil {
		return nil, err
	}

	return transformer.ToPost(p), nil
}

func (r *mutationResolver) DeletePost(ctx context.Context, input model.DeletePostRequest) (*model.Result, error) {
	result := new(model.Result)
	result.IsSuccess = false

	p, err := bmodel.GetPost(ctx, input.ID)
	if err != nil {
		return result, err
	}

	if p == nil {
		return result, fmt.Errorf("Post not found")
	}

	if err := bmodel.DeletePost(ctx, p); err != nil {
		return result, err
	}

	result.IsSuccess = true

	return result, nil
}

func (r *queryResolver) GetPostByID(ctx context.Context, id string) (*model.Post, error) {
	p, err := bmodel.GetPost(ctx, id)
	if err != nil {
		return nil, err
	}

	return transformer.ToPost(p), nil
}

func (r *queryResolver) GetPosts(ctx context.Context, cursor string, limit int) (*model.GetPostsResponse, error) {
	posts, newCursor, err := bmodel.ListPost(ctx, limit, cursor)

	if err != nil {
		return &model.GetPostsResponse{Posts: nil}, err
	}

	return &model.GetPostsResponse{Posts: transformer.ToPosts(posts), Cursor: newCursor}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
