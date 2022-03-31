package dataloaders

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/Ammce/hackernews/adapters/graph/models"
	"github.com/graph-gophers/dataloader"
)

type CommentMock struct {
	ID          string
	Text        string
	ArticleId   string
	CreatedById string
	CreatedAt   string
}

func CommentDataLoader(db *sql.DB) *dataloader.Loader {

	batchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		fmt.Println("Keys", keys)
		str := strings.Join(keys.Keys(), ", ")

		sqlStatement := fmt.Sprintf("SELECT c.id, c.text, c.article_id, c.created_by_id , c.created_at  AS comments FROM articles a JOIN comments c ON a.id = c.article_id WHERE a.id IN (%s)", str)
		row, err := db.Query(sqlStatement)
		if err != nil {
			return nil
		}
		defer row.Close()

		keysAndComments := map[string][]*models.Comment{}

		for row.Next() {
			var comment models.Comment
			err := row.Scan(&comment.ID, &comment.Text, &comment.ArticleId, &comment.CreatedById, &comment.CreatedAt)
			if err != nil {
				fmt.Println(err)
				return nil
			}
			keysAndComments[comment.ArticleId] = append(keysAndComments[comment.ArticleId], &comment)
		}

		fmt.Println(keysAndComments)

		defer row.Close()

		// var u = make(map[string]*models.Article, len(keys))

		// for _, art := range articlesSlice {
		// 	u[art.ID] = art
		// }

		var results = make([]*dataloader.Result, len(keys))

		for i, id := range keys {
			results[i] = &dataloader.Result{Data: keysAndComments[id.String()]}
		}
		return results
	}

	// create Loader with an in-memory cache
	loader := dataloader.NewBatchedLoader(batchFn)
	return loader
}
