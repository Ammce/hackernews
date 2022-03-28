package article

type ArticleService interface {
	// TODO - Change this to CreateNewsInput
	CreateArticle(*Article) (*Article, error)
	GetArticleById(ArticleId string) (*Article, error)
	GetAllArticles() ([]*Article, error)
}
