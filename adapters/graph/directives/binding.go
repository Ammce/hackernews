package directives

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/Ammce/hackernews/adapters/graph/generated"
	"github.com/go-playground/validator"
)

var (
	validate *validator.Validate
)

func init() {
	validate = validator.New()
}

func SetupBindingDirective(c *generated.Config) {
	c.Directives.Binding = func(ctx context.Context, obj interface{}, next graphql.Resolver, constraint string) (interface{}, error) {
		val, err := next(ctx)

		if err != nil {
			fmt.Println("Something really bad happened")
			return nil, err
		}
		fieldName := *graphql.GetPathContext(ctx).Field

		err = validate.Var(val, constraint)
		if err != nil {
			fmt.Println(err)
			validationErrors := err.(validator.ValidationErrors)
			transErr := fmt.Errorf("%s%+v", fieldName, validationErrors[0])
			return val, transErr
		}

		return val, nil
	}
}
