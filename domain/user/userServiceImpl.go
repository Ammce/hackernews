package user

import "fmt"

type UserServiceImpl struct {
	userRepo UserRepository
}

func (ur UserServiceImpl) CreateUser(user User) (*User, error) {
	fmt.Println("User je", user)
	return nil, nil
}
