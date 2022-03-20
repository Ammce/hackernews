package user

type UserService interface {
	CreateUser(*User) (*User, error)
}
