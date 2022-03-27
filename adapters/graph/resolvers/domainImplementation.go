package graph

import (
	authDomain "github.com/Ammce/hackernews/domain/auth"
	newsDomain "github.com/Ammce/hackernews/domain/news"
	userDomain "github.com/Ammce/hackernews/domain/user"
)

type DomainGraphQL struct {
	UserService userDomain.UserService
	AuthService authDomain.AuthService
	NewsService newsDomain.NewsService
}
