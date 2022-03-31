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
	"github.com/Ammce/hackernews/domain/article"
	"github.com/graph-gophers/dataloader"
)

func (r *articleResolver) CreatedBy(ctx context.Context, obj *models.Article) (*models.User, error) {
	thunk := r.UserDataLoader.Load(context.TODO(), dataloader.StringKey(obj.CreatedById)) // StringKey is a convenience method that make wraps string to implement `Key` interface
	result, err := thunk()
	if err != nil {
		fmt.Println("Erro se desio", err)
	}
	return result.(*models.User), nil
}

func (r *articleResolver) ApprovedBy(ctx context.Context, obj *models.Article) (*models.User, error) {
	thunk := r.UserDataLoader.Load(context.TODO(), dataloader.StringKey(obj.CreatedById)) // StringKey is a convenience method that make wraps string to implement `Key` interface
	result, err := thunk()
	if err != nil {
		fmt.Println("Erro se desio", err)
	}
	return result.(*models.User), nil
}

func (r *articleResolver) Comments(ctx context.Context, obj *models.Article) ([]*models.Comment, error) {
	// TODO - Handle comments here. Probably Data loader for comments
	thunk := r.CommentDataLoader.Load(context.TODO(), dataloader.StringKey(obj.ID)) // StringKey is a convenience method that make wraps string to implement `Key` interface
	// thunk := r.CommentDataLoader.Load(context.TODO(), dataloader.StringKey("10")) // StringKey is a convenience method that make wraps string to implement `Key` interface
	result, err := thunk()
	if err != nil {
		// fmt.Println("Erro se desio", err)
	}
	fmt.Println(result)
	return result.([]*models.Comment), nil
}

func (r *mutationResolver) CreateArticle(ctx context.Context, input inputs.ArticleInput) (*models.Article, error) {
	article, err := r.Domain.ArticleService.CreateArticle(mappers.ArticleInputToArticleDomain(&input))
	if err != nil {
		return nil, err
	}
	return mappers.ArticleDomainToArticleGraphQL(article), nil
}

func (r *queryResolver) Article(ctx context.Context, articleID string) (*models.Article, error) {
	article, errD := r.Domain.ArticleService.GetArticleById(articleID)
	if errD != nil {
		return nil, errors.New("error getting the domain news")
	}
	return mappers.ArticleDomainToArticleGraphQL(article), nil
}

func (r *queryResolver) Articles(ctx context.Context, filter *inputs.ArticleFilterInput) ([]*models.Article, error) {
	inputFilter := &article.ArticleFilter{}
	if filter != nil {
		inputFilter.CreatedById = filter.CreatedById
	}
	allDomainArticles, errD := r.Domain.ArticleService.GetAllArticles(inputFilter)

	if errD != nil {
		return nil, errors.New("error getting the domain news")
	}

	var allGqlArticles []*models.Article

	for _, da := range allDomainArticles {
		allGqlArticles = append(allGqlArticles, mappers.ArticleDomainToArticleGraphQL(da))
	}

	return allGqlArticles, nil
}

// Article returns generated.ArticleResolver implementation.
func (r *Resolver) Article() generated.ArticleResolver { return &articleResolver{r} }

type articleResolver struct{ *Resolver }
