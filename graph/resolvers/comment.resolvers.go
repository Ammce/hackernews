package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Ammce/hackernews/graph/generated"
	mocked_data "github.com/Ammce/hackernews/mock"
	"github.com/Ammce/hackernews/models"
)

func (r *commentResolver) CreatedBy(ctx context.Context, obj *models.Comment) (*models.User, error) {
	return &mocked_data.MockUser, nil
}

func (r *commentResolver) News(ctx context.Context, obj *models.Comment) (*models.News, error) {
	return &mocked_data.MockNews1, nil
}

func (r *queryResolver) Comment(ctx context.Context) (*models.Comment, error) {
	return &mocked_data.CommentMock1, nil
}

func (r *queryResolver) Comments(ctx context.Context) ([]*models.Comment, error) {
	return mocked_data.AllComments, nil
}

// Comment returns generated.CommentResolver implementation.
func (r *Resolver) Comment() generated.CommentResolver { return &commentResolver{r} }

type commentResolver struct{ *Resolver }