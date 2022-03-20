package user

import "fmt"

type UserServiceImpl struct {
	UserRepo UserRepository
}

func (ur UserServiceImpl) CreateUser(user User) (*User, error) {
	fmt.Println("User je", user)
	ur.UserRepo.SaveUser(user)
	return &User{
		Username: "Ammce",
		Email:    "amcenp@gmail.com",
		ID:       "001",
	}, nil
}
