package graph

import (
	"database/sql"

	"github.com/Ammce/hackernews/adapters/graph"
	"github.com/graph-gophers/dataloader"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
type Resolver struct {
	DB                *sql.DB
	Domain            graph.DomainGraphQL
	UserDataLoader    *dataloader.Loader
	ArticleDataLoader *dataloader.Loader
	CommentDataLoader *dataloader.Loader
}
