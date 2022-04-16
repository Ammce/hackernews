package externalarticle

type ExternalArticleService interface {
	getTopArticlesPerCountry(country *string) ([]*ExternalArticle, error)
}
