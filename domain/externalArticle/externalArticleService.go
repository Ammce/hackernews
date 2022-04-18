package externalarticle

type ExternalArticleService interface {
	GetTopExternalArticlesPerCountry(country *string) ([]*ExternalArticle, error)
	GetExternalArticlesByTopics(topics []string) ([]*ExternalArticlesByTopic, error)
}
