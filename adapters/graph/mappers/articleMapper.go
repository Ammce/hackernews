package mappers

import (
	"github.com/Ammce/hackernews/adapters/graph/models"
	"github.com/Ammce/hackernews/adapters/graph/models/inputs"
	"github.com/Ammce/hackernews/domain/article"
)

func ArticleInputToArticleDomain(articleInput *inputs.ArticleInput) *article.Article {
	return &article.Article{
		Title:       articleInput.Title,
		Text:        articleInput.Text,
		CreatedById: articleInput.CreatedById,
	}
}

func ArticleDomainToArticleGraphQL(newsDomain *article.Article) *models.Article {
	return &models.Article{
		ID:           newsDomain.ID,
		Title:        newsDomain.Title,
		Text:         newsDomain.Text,
		CreatedById:  newsDomain.CreatedById,
		CreatedAt:    newsDomain.CreatedAt,
		ApprovedById: newsDomain.ApprovedById,
	}
}
