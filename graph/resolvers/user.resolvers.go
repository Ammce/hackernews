package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Ammce/hackernews/graph/generated"
	mocked_data "github.com/Ammce/hackernews/mock"
	"github.com/Ammce/hackernews/models"
)

func (r *queryResolver) User(ctx context.Context) (*models.User, error) {
	return &mocked_data.MockUser, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	return []*models.User{&mocked_data.MockUser}, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
