package news

type NewsServiceImpl struct {
	NewsRepo NewsRepository
}

func (ns NewsServiceImpl) CreateNews(n *News) (*News, error) {
	return ns.NewsRepo.SaveNews(n)
}

func (ns NewsServiceImpl) GetNewsById(newsId string) (*News, error) {
	return ns.NewsRepo.GetNewsById(newsId)
}
func (ns NewsServiceImpl) GetAllNews() ([]*News, error) {
	return ns.NewsRepo.GetAllNews()
}

func NewNewsServiceImpl(newsRepo NewsRepository) NewsServiceImpl {
	return NewsServiceImpl{
		NewsRepo: newsRepo,
	}
}
