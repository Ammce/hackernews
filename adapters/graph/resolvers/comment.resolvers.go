package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/Ammce/hackernews/adapters/graph/generated"
	"github.com/Ammce/hackernews/adapters/graph/mappers"
	"github.com/Ammce/hackernews/adapters/graph/models"
	"github.com/Ammce/hackernews/adapters/graph/models/inputs"
	"github.com/graph-gophers/dataloader"
)

func (r *commentResolver) CreatedBy(ctx context.Context, obj *models.Comment) (*models.User, error) {
	thunk := r.UserDataLoader.Load(context.TODO(), dataloader.StringKey(obj.CreatedById)) // StringKey is a convenience method that make wraps string to implement `Key` interface
	result, err := thunk()
	if err != nil {
		fmt.Println("Erro se desio", err)
	}
	return result.(*models.User), nil
}

func (r *commentResolver) Article(ctx context.Context, obj *models.Comment) (*models.Article, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *commentResolver) IPAddress(ctx context.Context, obj *models.Comment) (string, error) {
	return "192.168.0.1", nil
}

func (r *mutationResolver) CreateComment(ctx context.Context, input inputs.CommentInput) (*models.Comment, error) {
	comment, errD := r.Domain.CommentService.CreateComment(mappers.CommentInputToCommentDomain(&input))
	if errD != nil {
		return nil, errD
	}
	return mappers.CommentDomainToCommentGraphQL(comment), nil
}

func (r *queryResolver) Comment(ctx context.Context, commentID string) (*models.Comment, error) {
	comment, errD := r.Domain.CommentService.GetCommentById(commentID)
	if errD != nil {
		return nil, errors.New("error getting the domain comments")
	}
	return mappers.CommentDomainToCommentGraphQL(comment), nil
}

func (r *queryResolver) Comments(ctx context.Context) ([]*models.Comment, error) {
	allDomainComments, errD := r.Domain.CommentService.GetAllComments()

	if errD != nil {
		return nil, errors.New("error getting the domain comments")
	}

	var allGqlComments []*models.Comment

	for _, dc := range allDomainComments {
		allGqlComments = append(allGqlComments, mappers.CommentDomainToCommentGraphQL(dc))
	}

	return allGqlComments, nil
}

// Comment returns generated.CommentResolver implementation.
func (r *Resolver) Comment() generated.CommentResolver { return &commentResolver{r} }

type commentResolver struct{ *Resolver }
