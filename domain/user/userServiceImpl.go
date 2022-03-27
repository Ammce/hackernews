package user

import (
	"errors"
)

type UserServiceImpl struct {
	UserRepo UserRepository
}

func (ur UserServiceImpl) CreateUser(user *User) (*User, error) {
	user.HashPassword(user.Password)
	savedUser, err := ur.UserRepo.SaveUser(user)
	if err != nil {
		return nil, errors.New("invalid query execution")
	}
	return savedUser, nil
}

func (ur UserServiceImpl) GetUser(userId string) (*User, error) {
	return ur.UserRepo.GetUserById(userId)
}

func (ur UserServiceImpl) GetUsers() ([]*User, error) {
	return ur.UserRepo.GetAllUsers()
}

func NewUserServiceImpl(userRepo UserRepository) UserServiceImpl {
	return UserServiceImpl{
		UserRepo: userRepo,
	}
}
