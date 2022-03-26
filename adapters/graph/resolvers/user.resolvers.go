package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/Ammce/hackernews/adapters/graph/helpers"
	"github.com/Ammce/hackernews/adapters/graph/mappers"
	"github.com/Ammce/hackernews/adapters/graph/models"
	"github.com/Ammce/hackernews/adapters/graph/models/inputs"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input *inputs.UserInput) (*models.User, error) {
	mappedUserInput := mappers.UserInputToUserDomain(input)
	domainUser, err := r.Domain.UserService.CreateUser(mappedUserInput)
	if err != nil {
		return nil, err
	}
	return mappers.UserDomainToUserGraphQL(domainUser), nil
}

func (r *mutationResolver) Login(ctx context.Context, input inputs.LoginInput) (*models.UserWithToken, error) {
	loginData, err := r.Domain.AuthService.Login(input.Email, input.Password)
	if err != nil {
		return nil, err
	}
	return &models.UserWithToken{
		Token: loginData.Token,
		User:  (*models.User)(loginData.User),
	}, nil
}

func (r *queryResolver) User(ctx context.Context, userID string) (*models.User, error) {
	user, err := r.Domain.UserService.GetUser(userID)
	if err != nil {
		return nil, err
	}
	return mappers.UserDomainToUserGraphQL(user), nil
}

func (r *queryResolver) Self(ctx context.Context) (*models.User, error) {
	userData := helpers.GetUsersDataFromContext(ctx)
	if userData == nil {
		return nil, errors.New("fIND ME -> No User data found")
	}

	user, err := r.Domain.UserService.GetUser(userData.UserId)
	if err != nil {
		return nil, err
	}
	return mappers.UserDomainToUserGraphQL(user), nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	usersDomain, err := r.Domain.UserService.GetUsers()
	if err != nil {
		return nil, errors.New("fIND ME -> Users , not found")
	}

	var usersGraphQL []*models.User

	for _, userDomain := range usersDomain {
		userGraphQL := mappers.UserDomainToUserGraphQL(userDomain)
		usersGraphQL = append(usersGraphQL, userGraphQL)

	}

	return usersGraphQL, nil

}
