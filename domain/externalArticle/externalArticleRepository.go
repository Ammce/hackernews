package externalarticle

type ExternalArticleRepository interface {
	GetTopExternalArticlesPerCountry(country *string) ([]*ExternalArticle, error)
	GetExternalArticlesByTopics(topics []string) ([]*ExternalArticlesByTopic, error)
}
