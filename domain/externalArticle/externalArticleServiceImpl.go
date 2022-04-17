package externalarticle

type ExternalArticleServiceImpl struct {
	ExternalArticleRepo ExternalArticleRepository
}

func (ea ExternalArticleServiceImpl) GetTopArticlesPerCountry(country *string) ([]*ExternalArticle, error) {
	topArticlesPerCountry, err := ea.ExternalArticleRepo.GetTopArticlesPerCountry(country)
	if err != nil {
		return nil, err
	}
	return topArticlesPerCountry, nil
}

func NewExternalArticleServiceImpl(ear ExternalArticleRepository) ExternalArticleServiceImpl {
	return ExternalArticleServiceImpl{ExternalArticleRepo: ear}
}
