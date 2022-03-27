package news

type NewsRepository interface {
	SaveNews(*News) (*News, error)
}
