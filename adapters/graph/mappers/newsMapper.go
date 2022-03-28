package mappers

import (
	"github.com/Ammce/hackernews/adapters/graph/models"
	"github.com/Ammce/hackernews/adapters/graph/models/inputs"
	"github.com/Ammce/hackernews/domain/news"
)

func NewsInputToNewsDomain(newsInput *inputs.NewsInput) *news.News {
	return &news.News{
		Title:       newsInput.Title,
		Text:        newsInput.Text,
		CreatedById: newsInput.CreatedById,
	}
}

func NewsDomainToNewsGraphQL(newsDomain *news.News) *models.News {
	return &models.News{
		ID:           newsDomain.ID,
		Title:        newsDomain.Title,
		Text:         newsDomain.Text,
		CreatedById:  newsDomain.CreatedById,
		CreatedAt:    newsDomain.CreatedAt,
		ApprovedById: newsDomain.ApprovedById,
	}
}
