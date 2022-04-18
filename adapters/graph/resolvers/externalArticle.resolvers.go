package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Ammce/hackernews/adapters/graph/mappers"
	"github.com/Ammce/hackernews/adapters/graph/models"
)

func (r *queryResolver) GetTopExternalArticlesByCountry(ctx context.Context, country *string) ([]*models.ExternalArticle, error) {
	resp, err := r.Domain.ExternalArticleService.GetTopArticlesPerCountry(country)
	if err != nil {
		return nil, err
	}

	var externalArticlesGraphQL []*models.ExternalArticle
	for _, ea := range resp {
		mappedEa := mappers.ExternalArticleDomainToExternalArticleGraphQL(ea)
		externalArticlesGraphQL = append(externalArticlesGraphQL, mappedEa)
	}

	return externalArticlesGraphQL, nil
}

func (r *queryResolver) GetExternalArticlesByTopics(ctx context.Context, topics []string) ([]*models.ExternalArticlesByTopic, error) {
	panic(fmt.Errorf("not implemented"))
}
