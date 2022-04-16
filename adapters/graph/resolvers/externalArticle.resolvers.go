package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Ammce/hackernews/adapters/graph/models"
)

func (r *queryResolver) GetTopArticlesPerCountry(ctx context.Context, country *string) ([]*models.ExternalArticle, error) {
	ct := "rs"
	resp, err := r.Domain.ExternalArticleService.GetTopArticlesPerCountry(&ct)
	if err != nil {
		return nil, err
	}
	fmt.Println("Response je", resp)
	return nil, nil
}
