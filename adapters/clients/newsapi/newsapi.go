package newsapi

import (
	"fmt"

	externalArticle "github.com/Ammce/hackernews/domain/externalArticle"
)

type NewsApi struct{}

func (n NewsApi) GetTopArticlesPerCountry(country *string) ([]*externalArticle.ExternalArticle, error) {
	fmt.Println("Hello")
	ea := externalArticle.ExternalArticle{
		Author: "Ammce",
	}
	return []*externalArticle.ExternalArticle{&ea}, nil
}

func NewNewsApi() NewsApi {
	return NewsApi{}
}
