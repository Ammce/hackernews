package repositories

import (
	"database/sql"
	"log"
	"strconv"

	"github.com/Ammce/hackernews/domain/user"
)

type UserRepositoryImpl struct {
	DB *sql.DB
}

func (ur UserRepositoryImpl) SaveUser(u *user.User) (*user.User, error) {
	// sqlStatement := `INSERT INTO users (email, username, password) VALUES ($1, $2, $3) RETURNING id`
	sqlStatement := `INSERT INTO users (email, username, password) VALUES ($1, $2, $3) RETURNING id`

	var id int64
	// err := ur.DB.QueryRow(sqlStatement, u.Email, u.Username, u.Password).Scan(&id)
	err := ur.DB.QueryRow(sqlStatement, u.Email, u.Username, u.Password).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	return &user.User{
		Username: u.Username,
		Email:    u.Email,
		ID:       strconv.FormatInt(id, 10),
		Password: u.Password,
	}, nil
}
