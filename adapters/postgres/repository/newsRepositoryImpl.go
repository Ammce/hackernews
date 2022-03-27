package repositories

import (
	"database/sql"
	"strconv"

	"github.com/Ammce/hackernews/domain/news"
)

type NewsRepositoryImpl struct {
	DB *sql.DB
}

func (nr NewsRepositoryImpl) SaveNews(n *news.News) (*news.News, error) {
	sqlStatement := `INSERT INTO news (title, text, created_by_id) VALUES ($1, $2, $3) returning id`

	var id int64

	if err := nr.DB.QueryRow(sqlStatement, n.Title, n.Text, n.CreatedById).Scan(&id); err != nil {
		return nil, err
	}
	return &news.News{
		ID:          strconv.FormatInt(id, 10),
		Text:        n.Text,
		Title:       n.Title,
		CreatedById: n.CreatedById,
	}, nil
}

func NewNewsRepositoryImpl(db *sql.DB) NewsRepositoryImpl {
	return NewsRepositoryImpl{
		DB: db,
	}
}
