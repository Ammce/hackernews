package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
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
	return mappers.NewsDomainToNewsGraphQ(news), nil
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
	return &mocked_data.MockNews1, nil
}

func (r *queryResolver) AllNews(ctx context.Context) ([]*models.News, error) {
	sqlStatement := `SELECT id, text, title, created_by_id FROM news`

	rows, err := r.DB.Query(sqlStatement)

	if err != nil {
		return nil, err
	}

	var allNews []*models.News

	for rows.Next() {
		var news models.News
		err := rows.Scan(&news.ID, &news.Text, &news.Title, &news.CreatedById)
		if err != nil {
			return nil, err
		}
		allNews = append(allNews, &news)
	}

	defer rows.Close()

	return allNews, nil
}

// News returns generated.NewsResolver implementation.
func (r *Resolver) News() generated.NewsResolver { return &newsResolver{r} }

type newsResolver struct{ *Resolver }
