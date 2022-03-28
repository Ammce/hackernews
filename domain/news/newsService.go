package news

type NewsService interface {
	// TODO - Change this to CreateNewsInput
	CreateNews(*News) (*News, error)
	GetNewsById(newsId string) (*News, error)
	GetAllNews() ([]*News, error)
}
