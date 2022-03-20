package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"log"

	"github.com/Ammce/hackernews/adapters/graph/mappers"
	"github.com/Ammce/hackernews/adapters/graph/models"
	"github.com/Ammce/hackernews/adapters/graph/models/inputs"
	"github.com/Ammce/hackernews/domain/user"
	mocked_data "github.com/Ammce/hackernews/mock"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input *inputs.UserInput) (*models.User, error) {
	mappedUserInput := mappers.UserInputToUserDomain(input)
	domainUser, _ := r.Domain.UserService.CreateUser(mappedUserInput)
	return mappers.UserDomainToUserGraphQL(domainUser), nil
}

func (r *queryResolver) User(ctx context.Context) (*models.User, error) {
	return &mocked_data.MockUser, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {

	r.Domain.UserService.CreateUser(user.User{ID: "123", Email: "amcenp", Username: "ammce"})

	sqlStatement := `SELECT id, username, email FROM users;`

	var users []*models.User

	rows, err := r.DB.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	for rows.Next() {
		var user models.User

		// unmarshal the row object to user
		err = rows.Scan(&user.ID, &user.Username, &user.Email)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the user in the users slice
		users = append(users, &user)

	}

	defer rows.Close()

	return users, nil
}
