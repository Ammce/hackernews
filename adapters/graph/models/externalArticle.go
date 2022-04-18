package models

type ExternalArticleSource struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ExternalArticle struct {
	Source      ExternalArticleSource `json:"source"`
	Author      string                `json:"author"`
	Description string                `json:"description"`
	Title       string                `json:"title"`
	Url         string                `json:"url"`
	UrlToImage  string                `json:"urlToImage"`
	PublishedAt string                `json:"publishedAt"`
	Content     string                `json:"content"`
}

type ExternalArticlesByTopic struct {
	Topic    string            `json:"topic"`
	Articles []ExternalArticle `json:"articles"`
}
