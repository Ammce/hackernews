package user

type UserRepository interface {
	SaveUser(*User) (*User, error)
	GetUserById(userId string) (*User, error)
	GetAllUsers() ([]*User, error)
}
