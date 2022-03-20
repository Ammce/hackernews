package user

type UserRepository interface {
	createUser(User) User
}
