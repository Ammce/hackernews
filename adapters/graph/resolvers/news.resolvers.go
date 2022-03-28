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
	mocked_data "github.com/Ammce/hackernews/mock"
	"github.com/graph-gophers/dataloader"
)

func (r *mutationResolver) CreateNews(ctx context.Context, input inputs.NewsInput) (*models.News, error) {
	news, err := r.Domain.NewsService.CreateNews(mappers.NewsInputToNewsDomain(&input))
	if err != nil {
		return nil, err
	}
	return mappers.NewsDomainToNewsGraphQL(news), nil
}

func (r *newsResolver) CreatedBy(ctx context.Context, obj *models.News) (*models.User, error) {
	thunk := r.UserDataLoader.Load(context.TODO(), dataloader.StringKey(obj.CreatedById)) // StringKey is a convenience method that make wraps string to implement `Key` interface
	result, err := thunk()
	if err != nil {
		fmt.Println("Erro se desio", err)
	}
	return result.(*models.User), nil
}

func (r *newsResolver) ApprovedBy(ctx context.Context, obj *models.News) (*models.User, error) {
	thunk := r.UserDataLoader.Load(context.TODO(), dataloader.StringKey(obj.CreatedById)) // StringKey is a convenience method that make wraps string to implement `Key` interface
	result, err := thunk()
	if err != nil {
		fmt.Println("Erro se desio", err)
	}
	return result.(*models.User), nil
}

func (r *newsResolver) Comments(ctx context.Context, obj *models.News) ([]*models.Comment, error) {
	// TODO - Handle comments here. Probably Data loader for comments
	var currentModels []*models.Comment
	for _, d := range mocked_data.AllComments {
		if d.NewsId == obj.ID {
			currentModels = append(currentModels, d)
		}
	}
	return currentModels, nil
}

func (r *queryResolver) News(ctx context.Context, newsID string) (*models.News, error) {
	news, errD := r.Domain.NewsService.GetNewsById(newsID)
	if errD != nil {
		return nil, errors.New("error getting the domain news")
	}
	return mappers.NewsDomainToNewsGraphQL(news), nil
}

func (r *queryResolver) AllNews(ctx context.Context) ([]*models.News, error) {

	allDomainNews, errD := r.Domain.NewsService.GetAllNews()
	if errD != nil {
		return nil, errors.New("error getting the domain news")
	}

	var allGqlNews []*models.News

	for _, dn := range allDomainNews {
		allGqlNews = append(allGqlNews, mappers.NewsDomainToNewsGraphQL(dn))
	}

	return allGqlNews, nil
}

// News returns generated.NewsResolver implementation.
func (r *Resolver) News() generated.NewsResolver { return &newsResolver{r} }

type newsResolver struct{ *Resolver }
