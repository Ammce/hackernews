package graph

import (
	"github.com/Ammce/hackernews/domain/auth"
	userDomain "github.com/Ammce/hackernews/domain/user"
)

type DomainGraphQL struct {
	UserService userDomain.UserService
	AuthService auth.AuthService
}
