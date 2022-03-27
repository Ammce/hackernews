package news

type NewsServiceImpl struct {
	NewsRepo NewsRepository
}

func (ns NewsServiceImpl) CreateNews(n *News) (*News, error) {
	return ns.NewsRepo.SaveNews(n)
}

func NewNewsServiceImpl(newsRepo NewsRepository) NewsServiceImpl {
	return NewsServiceImpl{
		NewsRepo: newsRepo,
	}
}
