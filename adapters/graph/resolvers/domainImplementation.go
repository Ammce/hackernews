package graph

import userDomain "github.com/Ammce/hackernews/domain/user"

type DomainGraphQL struct {
	userService userDomain.UserService
}

func DomainImplementation(us userDomain.UserService) DomainGraphQL {
	return DomainGraphQL{
		userService: us,
	}
}
