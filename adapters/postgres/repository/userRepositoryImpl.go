package repositories

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/Ammce/hackernews/domain/user"
)

type UserRepositoryImpl struct {
	DB *sql.DB
}

func (ur UserRepositoryImpl) SaveUser(u *user.User) (*user.User, error) {
	sqlStatement := `INSERT INTO users (email, username, password) VALUES ($1, $2, $3) RETURNING id`

	var id int64

	if err := ur.DB.QueryRow(sqlStatement, u.Email, u.Username, u.Password).Scan(&id); err != nil {
		fmt.Println("Unable to execute the query.", err)
		return nil, err
	}

	return &user.User{
		Username: u.Username,
		Email:    u.Email,
		ID:       strconv.FormatInt(id, 10),
		Password: u.Password,
	}, nil
}

func (ur UserRepositoryImpl) GetUserById(userId string) (*user.User, error) {
	sqlStatement := fmt.Sprintf(`SELECT id, email, username FROM users WHERE id = %v`, userId)

	var user user.User

	if err := ur.DB.QueryRow(sqlStatement).Scan(&user.ID, &user.Email, &user.Username); err != nil {
		fmt.Println("Unable to execute the query GetUserById.", err)
		return nil, err
	}

	return &user, nil
}

func (ur UserRepositoryImpl) GetUserByField(field string, fieldValue string) (*user.User, error) {
	sqlStatement := fmt.Sprintf(`SELECT id, email, username FROM users WHERE %s = %v`, field, fieldValue)
	fmt.Println("DELET THIS QUERY AFTER - ", sqlStatement)

	var user user.User

	if err := ur.DB.QueryRow(sqlStatement).Scan(&user.ID, &user.Email, &user.Username); err != nil {
		fmt.Println("Unable to execute the query GetUserById.", err)
		return nil, err
	}

	return &user, nil
}

func (ur UserRepositoryImpl) GetAllUsers() ([]*user.User, error) {
	sqlStatement := `SELECT id, username, email FROM users;`

	var users []*user.User

	rows, err := ur.DB.Query(sqlStatement)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	for rows.Next() {
		var user user.User

		// unmarshal the row object to user
		err = rows.Scan(&user.ID, &user.Username, &user.Email)

		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}

		// append the user in the users slice
		users = append(users, &user)

	}

	defer rows.Close()

	return users, nil
}
