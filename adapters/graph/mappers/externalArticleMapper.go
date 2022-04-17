package mappers

import (
	"github.com/Ammce/hackernews/adapters/graph/models"
	externalarticle "github.com/Ammce/hackernews/domain/externalArticle"
)

func ExternalArticleDomainToExternalArticleGraphQL(externalArticleDomain *externalarticle.ExternalArticle) *models.ExternalArticle {
	return &models.ExternalArticle{
		Author:      externalArticleDomain.Author,
		Source:      models.ExternalArticleSource(externalArticleDomain.Source),
		Url:         externalArticleDomain.Url,
		UrlToImage:  externalArticleDomain.UrlToImage,
		Title:       externalArticleDomain.Title,
		Description: externalArticleDomain.Description,
		Content:     externalArticleDomain.Content,
		PublishedAt: externalArticleDomain.PublishedAt,
	}
}
