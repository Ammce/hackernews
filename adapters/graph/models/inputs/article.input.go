package inputs

type ArticleInput struct {
	Title       string
	Text        string
	CreatedById string
}

type ArticleFilterInput struct {
	CreatedById string
}
