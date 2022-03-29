package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/Ammce/hackernews/domain/comment"
)

type CommentRepositoryImpl struct {
	DB *sql.DB
}

func (cr CommentRepositoryImpl) SaveComment(c *comment.Comment) (*comment.Comment, error) {
	sqlStatement := `INSERT INTO comments (text, created_by_id, article_id) VALUES ($1, $2, $3) returning id created_at`

	var id int64
	var createdAt string

	if err := cr.DB.QueryRow(sqlStatement, c.Text, c.CreatedById, c.ArticleId).Scan(&id, &createdAt); err != nil {
		return nil, err
	}
	return &comment.Comment{
		ID:          strconv.FormatInt(id, 10),
		Text:        c.Text,
		CreatedById: c.CreatedById,
		ArticleId:   c.ArticleId,
		CreatedAt:   c.CreatedAt,
	}, nil
}

func (cr CommentRepositoryImpl) GetCommentById(commentId string) (*comment.Comment, error) {
	sqlStatement := fmt.Sprintf(`SELECT id, text, created_by_id, article_id, created_at FROM comments WHERE id = %s`, commentId)

	var comment comment.Comment

	if err := cr.DB.QueryRow(sqlStatement).Scan(&comment.ID, &comment.Text, &comment.CreatedById, &comment.ArticleId, &comment.CreatedAt); err != nil {
		return nil, err
	}

	return &comment, nil
}

func (cr CommentRepositoryImpl) GetAllComments() ([]*comment.Comment, error) {

	sqlStatement := `SELECT id, text, created_by_id, article_id, created_at FROM comments`

	var comments []*comment.Comment

	rows, err := cr.DB.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	for rows.Next() {
		var c comment.Comment

		err = rows.Scan(&c.ID, &c.Text, &c.CreatedById, &c.ArticleId, &c.CreatedAt)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		comments = append(comments, &c)

	}

	defer rows.Close()

	return comments, nil
}

func NewCommentRepositoryImpl(db *sql.DB) CommentRepositoryImpl {
	return CommentRepositoryImpl{
		DB: db,
	}
}
