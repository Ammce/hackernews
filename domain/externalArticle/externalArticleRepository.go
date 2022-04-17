package externalarticle

type ExternalArticleRepository interface {
	GetTopArticlesPerCountry(country *string) ([]*ExternalArticle, error)
}
