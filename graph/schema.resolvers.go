package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Ammce/hackernews/graph/generated"
	"github.com/Ammce/hackernews/graph/model"
	"github.com/Ammce/hackernews/models"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*models.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todos(ctx context.Context) ([]*models.Todo, error) {
	return []*models.Todo{
		{
			ID:   "abc-123",
			Text: "hello",
			Done: false,
		},
	}, nil
}

func (r *todoResolver) User(ctx context.Context, obj *models.Todo) (*models.User, error) {
	fmt.Println(obj)
	return &models.User{
		Name: "Amel",
		ID:   "224",
	}, nil
}

func (r *todoResolver) Comments(ctx context.Context, obj *models.Todo) ([]*models.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) Todos(ctx context.Context, obj *models.User) ([]*models.Todo, error) {
	return []*models.Todo{
		{
			ID:   "abc-224",
			Text: "hello",
			Done: false,
		},
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type todoResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
