package models

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"naemailme"`
	Password string `json:"password"`
}

type UserWithToken struct {
	User  *User
	Token string
}
