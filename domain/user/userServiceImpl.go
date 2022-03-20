package user

import "fmt"

type UserServiceImpl struct {
	userRepo UserRepository
}

func (ur UserServiceImpl) CreateUser(user User) (*User, error) {
	fmt.Println("User je", user)
	return &User{
		Username: "Ammce",
		Email:    "amcenp@gmail.com",
		ID:       "001",
	}, nil
}
