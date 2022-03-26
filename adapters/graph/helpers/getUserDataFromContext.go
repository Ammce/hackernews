package helpers

import (
	"context"

	"github.com/Ammce/hackernews/adapters/graph/middleware"
)

func GetUsersDataFromContext(ctx context.Context) *middleware.UserIDAndRoles {
	token := ctx.Value(middleware.AuthorizationHeader)
	userIdAndRoles := middleware.VerifyToken(token.(string))
	return userIdAndRoles
}
