package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Ammce/hackernews/graph/generated"
	mocked_data "github.com/Ammce/hackernews/mock"
	"github.com/Ammce/hackernews/models"
)

func (r *newsResolver) CreatedBy(ctx context.Context, obj *models.News) (*models.User, error) {
	return &mocked_data.MockUser, nil
}

func (r *newsResolver) ApprovedBy(ctx context.Context, obj *models.News) (*models.User, error) {
	return &mocked_data.MockUser, nil
}

func (r *newsResolver) Comments(ctx context.Context, obj *models.News) ([]*models.Comment, error) {
	var currentModels []*models.Comment
	for _, d := range mocked_data.AllComments {
		if d.NewsId == obj.ID {
			currentModels = append(currentModels, d)
		}
	}
	return currentModels, nil
}

func (r *queryResolver) News(ctx context.Context) (*models.News, error) {
	return &mocked_data.MockNews1, nil
}

func (r *queryResolver) AllNews(ctx context.Context) ([]*models.News, error) {
	return mocked_data.News, nil
}

// News returns generated.NewsResolver implementation.
func (r *Resolver) News() generated.NewsResolver { return &newsResolver{r} }

type newsResolver struct{ *Resolver }
