package graph

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/Ammce/hackernews/adapters/graph/models"
	"github.com/graph-gophers/dataloader"
)

func UserDataLoader(db *sql.DB) *dataloader.Loader {

	batchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		str := strings.Join(keys.Keys(), ", ")

		fmt.Println(str)
		sqlStatement := fmt.Sprintf("SELECT id, username, email FROM users WHERE id IN (%s)", str)
		row, err := db.Query(sqlStatement)
		if err != nil {
			return nil
		}

		var usersSlice []*models.User

		for row.Next() {
			var user models.User
			err := row.Scan(&user.ID, &user.Username, &user.Email)
			if err != nil {
				return nil
			}
			usersSlice = append(usersSlice, &user)
		}

		defer row.Close()

		var u = make(map[string]*models.User, len(keys))

		for _, usr := range usersSlice {
			u[usr.ID] = usr
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
