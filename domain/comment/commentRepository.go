package comment

type CommentRepository interface {
	SaveComment(*Comment) (*Comment, error)
	GetCommentById(string) (*Comment, error)
	GetAllComments() ([]*Comment, error)
}
