package directives

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/Ammce/hackernews/adapters/graph/generated"
	"github.com/Ammce/hackernews/adapters/graph/middleware"
)

func SetupIsAuthDirective(c *generated.Config) {
	c.Directives.IsAuth = func(ctx context.Context, obj interface{}, next graphql.Resolver) (interface{}, error) {
		token := ctx.Value(middleware.AuthorizationHeader)
		userIdAndRoles := middleware.VerifyToken(token.(string))
		if userIdAndRoles != nil {
			ctx = context.WithValue(ctx, middleware.UserDataKey, userIdAndRoles)
		}

		return next(ctx)
	}
}
