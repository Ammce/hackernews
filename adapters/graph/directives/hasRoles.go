package directives

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/Ammce/hackernews/adapters/graph/generated"
	"github.com/Ammce/hackernews/adapters/graph/middleware"
	"github.com/Ammce/hackernews/adapters/graph/models"
)

func SetupHasRolesDirective(c *generated.Config) {
	c.Directives.HasRoles = func(ctx context.Context, obj interface{}, next graphql.Resolver, roles []models.Role) (interface{}, error) {
		userData := ctx.Value(middleware.UserDataKey)
		if userData == nil {
			return nil, errors.New("calling @hasRoles directive before or without @isAuth directive")
		}
		castedUserData := userData.(*middleware.UserIDAndRoles)
		userRolesMap := map[string]bool{}
		for _, userRole := range castedUserData.Roles {
			userRolesMap[string(userRole)] = true
		}

		for _, requiredRole := range roles {
			if !userRolesMap[string(requiredRole)] {
				return nil, errors.New("no permissiosn for this")
			}
		}

		return next(ctx)
	}
}
