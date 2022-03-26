package auth

import "github.com/Ammce/hackernews/domain/user"

type UserWithToken struct {
	User  *user.User
	Token string
}
