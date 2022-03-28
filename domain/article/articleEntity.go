package article

type Article struct {
	ID           string
	Title        string
	Text         string
	Published    bool
	CreatedById  string
	CreatedAt    string
	ApprovedById string
}

type ArticleFilter struct {
	CreatedById string
}
