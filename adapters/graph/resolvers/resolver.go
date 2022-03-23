package graph

import (
	"database/sql"

	"github.com/graph-gophers/dataloader"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
type Resolver struct {
	DB             *sql.DB
	Domain         DomainGraphQL
	UserDataLoader *dataloader.Loader
}
