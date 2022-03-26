package user

type UserService interface {
	CreateUser(*User) (*User, error)
	GetUser(userId string) (*User, error)
	GetUsers() ([]*User, error)
}
