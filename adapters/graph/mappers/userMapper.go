package mappers

import (
	"github.com/Ammce/hackernews/adapters/graph/models"
	"github.com/Ammce/hackernews/adapters/graph/models/inputs"
	"github.com/Ammce/hackernews/domain/user"
)

func UserInputToUserDomain(userInput *inputs.UserInput) user.User {
	return user.User{
		Email:    userInput.Email,
		Username: userInput.Username,
		Password: userInput.Password,
	}
}

func UserDomainToUserGraphQL(userDomain *user.User) *models.User {
	return &models.User{
		ID:       userDomain.ID,
		Username: userDomain.Username,
		Email:    userDomain.Email,
	}
}
