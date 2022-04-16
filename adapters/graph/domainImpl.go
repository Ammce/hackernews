package graph

import (
	articleDomain "github.com/Ammce/hackernews/domain/article"
	authDomain "github.com/Ammce/hackernews/domain/auth"
	commentDomain "github.com/Ammce/hackernews/domain/comment"
	externalArticleDomain "github.com/Ammce/hackernews/domain/externalArticle"
	userDomain "github.com/Ammce/hackernews/domain/user"
)

type DomainGraphQL struct {
	UserService            userDomain.UserService
	AuthService            authDomain.AuthService
	ArticleService         articleDomain.ArticleService
	CommentService         commentDomain.CommentService
	ExternalArticleService externalArticleDomain.ExternalArticleService
}
