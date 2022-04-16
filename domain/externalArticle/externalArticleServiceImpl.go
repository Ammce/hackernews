package externalarticle

type ExternalArticleServiceImpl struct{}

func (ea ExternalArticleServiceImpl) GetTopArticlesPerCountry(country *string) ([]*ExternalArticle, error) {
	return nil, nil
}

func NewExternalArticleServiceImpl() ExternalArticleServiceImpl {
	return ExternalArticleServiceImpl{}
}
