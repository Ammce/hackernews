package news

type NewsRepository interface {
	SaveNews(*News) (*News, error)
	GetNewsById(string) (*News, error)
	GetAllNews() ([]*News, error)
}
