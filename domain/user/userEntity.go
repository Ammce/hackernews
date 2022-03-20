package user

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string
	Username string
	Email    string
	Password string
}

func (u *User) HashPassword(password string) error {
	passwordByte := []byte(password)
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordByte, bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error while hashing the password")
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) ComparePassword(password string) error {
	bytePassword := []byte(password)
	byteHashedPassword := []byte(u.Password)
	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}
