package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/Ammce/hackernews/domain/article"
)

type ArticleRepositoryImpl struct {
	DB *sql.DB
}

func (ar ArticleRepositoryImpl) SaveArticle(n *article.Article) (*article.Article, error) {
	sqlStatement := `INSERT INTO articles (title, text, created_by_id) VALUES ($1, $2, $3) returning id`

	var id int64

	if err := ar.DB.QueryRow(sqlStatement, n.Title, n.Text, n.CreatedById).Scan(&id); err != nil {
		return nil, err
	}
	return &article.Article{
		ID:          strconv.FormatInt(id, 10),
		Text:        n.Text,
		Title:       n.Title,
		CreatedById: n.CreatedById,
	}, nil
}

func (ar ArticleRepositoryImpl) GetArticleById(articleId string) (*article.Article, error) {
	sqlStatement := fmt.Sprintf(`SELECT id, title, text, created_by_id, created_at FROM articles WHERE id = %s`, articleId)

	var article article.Article

	if err := ar.DB.QueryRow(sqlStatement).Scan(&article.ID, &article.Title, &article.Text, &article.CreatedById, &article.CreatedAt); err != nil {
		return nil, err
	}

	return &article, nil
}

func (ar ArticleRepositoryImpl) GetAllArticles() ([]*article.Article, error) {
	sqlStatement := `SELECT id, title, text, created_by_id, created_at FROM articles`

	var articles []*article.Article

	rows, err := ar.DB.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	for rows.Next() {
		var a article.Article

		err = rows.Scan(&a.ID, &a.Title, &a.Text, &a.CreatedById, &a.CreatedAt)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		articles = append(articles, &a)

	}

	defer rows.Close()

	return articles, nil
}

func NewArticleRepositoryImpl(db *sql.DB) ArticleRepositoryImpl {
	return ArticleRepositoryImpl{
		DB: db,
	}
}
