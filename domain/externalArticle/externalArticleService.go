package externalarticle

type ExternalArticleService interface {
	GetTopArticlesPerCountry(country *string) ([]*ExternalArticle, error)
}
