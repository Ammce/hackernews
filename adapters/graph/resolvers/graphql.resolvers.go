package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Ammce/hackernews/adapters/graph/generated"
)

func (r *mutationResolver) HealtcheckMutation(ctx context.Context, str string) (string, error) {
	return str, nil
}

func (r *queryResolver) Healthcheck(ctx context.Context) (string, error) {
	return "Works just fine", nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
