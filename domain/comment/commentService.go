package comment

type CommentService interface {
	CreateComment(*Comment) (*Comment, error)
	GetCommentById(string) (*Comment, error)
	GetAllComments() ([]*Comment, error)
}
