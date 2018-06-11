package models

import (
	"github.com/satori/go.uuid"
	"database/sql"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

const (
	createUser = "INSERT INTO user (name, login, password) VALUES ($1, $2, $3)"
	getUser = "SELECT * FROM user WHERE id = $1"
	deleteUser = "DELETE FROM user WHERE id = $1"
	getUsers = "SELECT * FROM user"
)

type User struct {
	ID       uuid.UUID
	Name     string
	Login    string
	Password string
}

//will be deleteted!
var db *sql.DB

func init() {
	db, _, _ = sqlmock.New()
}

func CreateUser(user User) error {
	_, err := db.Exec(createUser, user.Name, user.Login, user.Password)

	if err != nil {
		return err
	}
	return nil
}

func GetUser(id uuid.UUID) (User, error) {
	user := User{}
	err := db.QueryRow(getUser, id).Scan(&user.ID, &user.Name, &user.Login, &user.Password)

	if err != nil {
		return user, err
	}
	return user, nil
}

func UpdateUser() {

}

func DeleteUser(id uuid.UUID) error {
	_, err := db.Exec(deleteUser, id)

	if err != nil {
		return err
	}
	return nil
}

func GetUsers() ([]User, error) {

	rows, err := db.Query(getUsers)
	if err != nil {
		return []User{}, err
	}

	users := make([]User, 0)

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Login, &u.Password); err != nil {
			return []User{}, err
		}
		users = append(users, u)
	}
	return users, nil
}