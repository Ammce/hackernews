package externalarticle

type ExternalArticleServiceImpl struct {
	ExternalArticleRepo ExternalArticleRepository
}

func (ea ExternalArticleServiceImpl) GetTopExternalArticlesPerCountry(country *string) ([]*ExternalArticle, error) {
	topExternalArticlesPerCountry, err := ea.ExternalArticleRepo.GetTopExternalArticlesPerCountry(country)
	if err != nil {
		return nil, err
	}
	return topExternalArticlesPerCountry, nil
}

func (ea ExternalArticleServiceImpl) GetExternalArticlesByTopics(topics []string) ([]*ExternalArticlesByTopic, error) {
	externalArticlesByTopics, err := ea.ExternalArticleRepo.GetExternalArticlesByTopics(topics)
	if err != nil {
		return nil, err
	}
	return externalArticlesByTopics, nil
}

func NewExternalArticleServiceImpl(ear ExternalArticleRepository) ExternalArticleServiceImpl {
	return ExternalArticleServiceImpl{ExternalArticleRepo: ear}
}
