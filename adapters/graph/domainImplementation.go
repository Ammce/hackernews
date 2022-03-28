package graph

import (
	articleDomain "github.com/Ammce/hackernews/domain/article"
	authDomain "github.com/Ammce/hackernews/domain/auth"
	userDomain "github.com/Ammce/hackernews/domain/user"
)

type DomainGraphQL struct {
	UserService    userDomain.UserService
	AuthService    authDomain.AuthService
	ArticleService articleDomain.ArticleService
}
