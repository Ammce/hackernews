package dataloaders

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/Ammce/hackernews/adapters/graph/models"
	"github.com/graph-gophers/dataloader"
)

func ArticleDataLoader(db *sql.DB) *dataloader.Loader {

	batchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		str := strings.Join(keys.Keys(), ", ")

		sqlStatement := fmt.Sprintf("SELECT id, title, text, created_by_id, created_at FROM articles WHERE id IN (%s)", str)
		row, err := db.Query(sqlStatement)
		if err != nil {
			return nil
		}

		var articlesSlice []*models.Article

		for row.Next() {
			var article models.Article
			err := row.Scan(&article.ID, &article.Title, &article.Text, &article.CreatedById, &article.CreatedAt)
			if err != nil {
				return nil
			}
			articlesSlice = append(articlesSlice, &article)
		}

		defer row.Close()

		var u = make(map[string]*models.Article, len(keys))

		for _, art := range articlesSlice {
			u[art.ID] = art
		}

		var results = make([]*dataloader.Result, len(keys))

		for i, id := range keys {
			results[i] = &dataloader.Result{Data: u[id.String()]}
		}
		return results
	}

	// create Loader with an in-memory cache
	loader := dataloader.NewBatchedLoader(batchFn)
	return loader
}
