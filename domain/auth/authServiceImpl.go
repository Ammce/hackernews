package auth

import (
	"errors"
	"fmt"

	"github.com/Ammce/hackernews/adapters/graph/middleware"
	"github.com/Ammce/hackernews/domain/user"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	UserRepo user.UserRepository
}

func (au AuthServiceImpl) Login(email string, password string) (*UserWithToken, error) {
	user, err := au.UserRepo.GetUserByField("email", email)
	if err != nil {
		return nil, errors.New("invalid login credentials - email")
	}

	hashErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if hashErr != nil {
		fmt.Println(hashErr)
		return nil, errors.New("invalid data pass")
	}
	signedToken := middleware.SignToken(user.ID)

	return &UserWithToken{
		Token: *signedToken,
		User:  user,
	}, nil
}

func NewAuthServiceImpl(userRepo user.UserRepository) AuthServiceImpl {
	return AuthServiceImpl{UserRepo: userRepo}
}
