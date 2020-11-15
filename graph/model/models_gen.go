// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Comment struct {
	ID        string `json:"id"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdAt"`
}

type CreatePostRequest struct {
	ID      *string `json:"id"`
	Title   string  `json:"title"`
	Content string  `json:"content"`
}

type DeletePostRequest struct {
	ID string `json:"id"`
}

type GetPostsResponse struct {
	Posts  []*Post `json:"posts"`
	Cursor string  `json:"cursor"`
}

type Post struct {
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	Content   string     `json:"content"`
	Comments  []*Comment `json:"comments"`
	CreatedAt string     `json:"createdAt"`
	UpdatedAt string     `json:"updatedAt"`
	DeletedAt *string    `json:"deletedAt"`
}

type Result struct {
	IsSuccess bool `json:"isSuccess"`
}

type UpdatePostRequest struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
