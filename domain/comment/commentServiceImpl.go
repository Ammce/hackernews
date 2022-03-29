package comment

type CommentServiceImpl struct {
	CommentRepo CommentRepository
}

func (csi CommentServiceImpl) CreateComment(comment *Comment) (*Comment, error) {
	return csi.CommentRepo.SaveComment(comment)
}

func (csi CommentServiceImpl) GetCommentById(commentId string) (*Comment, error) {
	return csi.CommentRepo.GetCommentById(commentId)
}

func (csi CommentServiceImpl) GetAllComments() ([]*Comment, error) {
	return csi.CommentRepo.GetAllComments()
}

func NewCommentServiceImpl(repo CommentRepository) CommentServiceImpl {
	return CommentServiceImpl{
		CommentRepo: repo,
	}
}
