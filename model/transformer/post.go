package transformer

import (
	"time"

	"github.com/wilsontwm/graphql-tutorial/graph/model"
	bmodel "github.com/wilsontwm/graphql-tutorial/model"
)

// ToPost :
func ToPost(i *bmodel.Post) *model.Post {
	o := new(model.Post)
	o.ID = i.ID.Hex()
	o.Title = i.Title
	o.Content = i.Content
	o.CreatedAt = i.CreatedAt.Format(time.RFC3339)
	o.UpdatedAt = i.UpdatedAt.Format(time.RFC3339)

	return o
}

// ToPosts :
func ToPosts(posts []*bmodel.Post) []*model.Post {
	var result []*model.Post
	for _, post := range posts {
		result = append(result, ToPost(post))
	}

	return result
}
