package externalarticle

type ExternalArticleSource struct {
	ID   string
	Name string
}

type ExternalArticle struct {
	Source      ExternalArticleSource
	Author      string
	Description string
	Title       string
	Url         string
	UrlToImage  string
	PublishedAt string
	Content     string
}
