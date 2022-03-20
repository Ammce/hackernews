package user

type UserRepository interface {
	SaveUser(*User) (*User, error)
}
