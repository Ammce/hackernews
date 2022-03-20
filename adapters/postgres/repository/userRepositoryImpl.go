package repositories

import (
	"database/sql"
	"fmt"

	"github.com/Ammce/hackernews/domain/user"
)

type UserRepositoryImpl struct {
	DB *sql.DB
}

func (ur UserRepositoryImpl) SaveUser(u user.User) user.User {
	fmt.Println("Trying to save user...")
	return user.User{
		Username: "Ammce",
		Email:    "amcenp@gmail.com",
		ID:       "001",
	}
}
