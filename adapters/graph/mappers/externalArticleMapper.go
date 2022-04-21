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

func ExternalArticlesDomainToExternalArticlesGraphQL(externalArticlesDomain []*externalarticle.ExternalArticle) []*models.ExternalArticle {
	var externalArticlesGraphQL []*models.ExternalArticle
	for _, ea := range externalArticlesDomain {
		mappedEa := ExternalArticleDomainToExternalArticleGraphQL(ea)
		externalArticlesGraphQL = append(externalArticlesGraphQL, mappedEa)
	}
	return externalArticlesGraphQL
}

func ExternalArticlesByTopicDomainToExternalArticleByTopicGraphQL(externalArticlesByTopicDomain []*externalarticle.ExternalArticlesByTopic) []*models.ExternalArticlesByTopic {

	var externalArticlesByTopicGraphQL []*models.ExternalArticlesByTopic

	for _, eaDomain := range externalArticlesByTopicDomain {
		var externalArticles []models.ExternalArticle
		for _, eaDomainArticle := range eaDomain.Articles {
			art := models.ExternalArticle{
				Title:       eaDomainArticle.Title,
				Description: eaDomainArticle.Description,
				Author:      eaDomainArticle.Author,
				Url:         eaDomainArticle.Url,
				UrlToImage:  eaDomainArticle.UrlToImage,
				Source:      models.ExternalArticleSource(eaDomainArticle.Source),
				PublishedAt: eaDomainArticle.PublishedAt,
			}
			externalArticles = append(externalArticles, art)
		}
		ea := models.ExternalArticlesByTopic{
			Topic:    eaDomain.Topic,
			Articles: externalArticles,
		}
		externalArticlesByTopicGraphQL = append(externalArticlesByTopicGraphQL, &ea)
	}

	return externalArticlesByTopicGraphQL
}
