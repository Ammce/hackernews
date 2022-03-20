package user

type UserServiceImpl struct {
	UserRepo UserRepository
}

func (ur UserServiceImpl) CreateUser(user *User) (*User, error) {
	savedUser, err := ur.UserRepo.SaveUser(user)
	if err != nil {
		return nil, err
	}
	return savedUser, nil
}
