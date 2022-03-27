package directives

import (
	"github.com/Ammce/hackernews/adapters/graph/generated"
)

func SetupDirectives(c *generated.Config) {
	SetupHasRolesDirective(c)
	SetupIsAuthDirective(c)
	SetupBindingDirective(c)
}
