package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/Ammce/hackernews/adapters/graph/mappers"
	"github.com/Ammce/hackernews/adapters/graph/models"
)

func (r *queryResolver) GetTopExternalArticlesByCountry(ctx context.Context, country *string) ([]*models.ExternalArticle, error) {
	resp, err := r.Domain.ExternalArticleService.GetTopExternalArticlesPerCountry(country)
	if err != nil {
		return nil, err
	}

	return mappers.ExternalArticlesDomainToExternalArticlesGraphQL(resp), nil
}

func (r *queryResolver) GetExternalArticlesByTopics(ctx context.Context, topics []string) ([]*models.ExternalArticlesByTopic, error) {
	externalArticlesByTopicDomain, err := r.Domain.ExternalArticleService.GetExternalArticlesByTopics(topics)
	externalArticlesByTopicGraphQL := mappers.ExternalArticlesByTopicDomainToExternalArticleByTopicGraphQL(externalArticlesByTopicDomain)
	if err != nil {
		return nil, err
	}
	return externalArticlesByTopicGraphQL, nil
}
