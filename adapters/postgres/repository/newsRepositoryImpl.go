package repositories

import (
	"database/sql"
	"fmt"
	"log"
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

func (nr NewsRepositoryImpl) GetNewsById(newsId string) (*news.News, error) {
	sqlStatement := fmt.Sprintf(`SELECT id, title, text, created_by_id, created_at FROM news WHERE id = %s`, newsId)

	var news news.News

	if err := nr.DB.QueryRow(sqlStatement).Scan(&news.ID, &news.Title, &news.Text, &news.CreatedById, &news.CreatedAt); err != nil {
		return nil, err
	}

	return &news, nil
}

func (nr NewsRepositoryImpl) GetAllNews() ([]*news.News, error) {
	sqlStatement := `SELECT id, title, text, created_by_id, created_at FROM news`

	var allNews []*news.News

	rows, err := nr.DB.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	for rows.Next() {
		var n news.News

		err = rows.Scan(&n.ID, &n.Title, &n.Text, &n.CreatedById, &n.CreatedAt)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		allNews = append(allNews, &n)

	}

	defer rows.Close()

	return allNews, nil
}

func NewNewsRepositoryImpl(db *sql.DB) NewsRepositoryImpl {
	return NewsRepositoryImpl{
		DB: db,
	}
}
