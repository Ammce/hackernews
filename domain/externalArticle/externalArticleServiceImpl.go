package externalarticle

type ExternalArticleServiceImpl struct{}

func (ea ExternalArticleServiceImpl) getTopArticlesPerCountry(country *string) ([]*ExternalArticle, error) {
	return nil, nil
}

func NewExternalArticleServiceImpl() ExternalArticleServiceImpl {
	return ExternalArticleServiceImpl{}
}
