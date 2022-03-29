package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Ammce/hackernews/adapters/graph/generated"
	"github.com/Ammce/hackernews/adapters/graph/models"
	"github.com/Ammce/hackernews/adapters/graph/models/inputs"
	mocked_data "github.com/Ammce/hackernews/mock"
)

func (r *commentResolver) CreatedBy(ctx context.Context, obj *models.Comment) (*models.User, error) {
	return &mocked_data.MockUser, nil
}

func (r *commentResolver) Article(ctx context.Context, obj *models.Comment) (*models.Article, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *commentResolver) IPAddress(ctx context.Context, obj *models.Comment) (string, error) {
	return "192.168.0.1", nil
}

func (r *mutationResolver) CreateComment(ctx context.Context, input inputs.CommentInput) (*models.Comment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Comment(ctx context.Context, commentID string) (*models.Comment, error) {
	return &mocked_data.CommentMock1, nil
}

func (r *queryResolver) Comments(ctx context.Context) ([]*models.Comment, error) {
	return mocked_data.AllComments, nil
}

// Comment returns generated.CommentResolver implementation.
func (r *Resolver) Comment() generated.CommentResolver { return &commentResolver{r} }

type commentResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *commentResolver) CreatedAt(ctx context.Context, obj *models.Comment) (string, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *commentResolver) ArticleID(ctx context.Context, obj *models.Comment) (string, error) {
	panic(fmt.Errorf("not implemented"))
}
