package article

type ArticleServiceImpl struct {
	ArticleRepo ArticleRepository
}

func (as ArticleServiceImpl) CreateArticle(a *Article) (*Article, error) {
	return as.ArticleRepo.SaveArticle(a)
}

func (as ArticleServiceImpl) GetArticleById(articleId string) (*Article, error) {
	return as.ArticleRepo.GetArticleById(articleId)
}
func (as ArticleServiceImpl) GetAllArticles(articleFilter *ArticleFilter) ([]*Article, error) {
	return as.ArticleRepo.GetAllArticles(articleFilter)
}

func NewArticleServiceImpl(articleRepostiroy ArticleRepository) ArticleServiceImpl {
	return ArticleServiceImpl{
		ArticleRepo: articleRepostiroy,
	}
}
