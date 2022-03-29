package mappers

import (
	"github.com/Ammce/hackernews/adapters/graph/models"
	"github.com/Ammce/hackernews/adapters/graph/models/inputs"
	"github.com/Ammce/hackernews/domain/comment"
)

func CommentInputToCommentDomain(commentInput *inputs.CommentInput) *comment.Comment {
	return &comment.Comment{
		Text:        commentInput.Text,
		CreatedById: commentInput.CreatedById,
		ArticleId:   commentInput.ArticleId,
	}
}

func CommentDomainToCommentGraphQL(commentDomain *comment.Comment) *models.Comment {
	return &models.Comment{
		ID:          commentDomain.ID,
		Text:        commentDomain.Text,
		CreatedById: commentDomain.CreatedById,
		CreatedAt:   commentDomain.CreatedAt,
		ArticleId:   commentDomain.ArticleId,
	}
}
