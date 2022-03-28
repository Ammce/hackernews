package article

type ArticleRepository interface {
	SaveArticle(*Article) (*Article, error)
	GetArticleById(string) (*Article, error)
	GetAllArticles() ([]*Article, error)
}
