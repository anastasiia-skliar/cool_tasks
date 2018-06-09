package models

import (
	"github.com/satori/go.uuid"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

type User struct {
	ID       uuid.UUID
	Name     string
	Login    string
	Password string
}

var DB, Mock, Err = sqlmock.New()

func CreateUser(user User) error {
	_, err := DB.Exec("INSERT INTO User (name, login, password) VALUES ($1, $1, $3)", user.Name, user.Login, user.Password)

	if err != nil {
		return err
	}
	return nil
}

func GetUser(id uuid.UUID) (*User, error) {
	user := &User{}
	err := DB.QueryRow("SELECT * FROM User WHERE id = $1", id).Scan(&user.Name, &user.Login, &user.Password)

	if err != nil {
		return user, err
	}
	return user, nil
}

func UpdateUser() {

}

func DeleteUser(id uuid.UUID) error {
	_, err := DB.Exec("DELETE FROM User WHERE id = $1", id)

	if err != nil {
		return err
	}
	return nil
}

func GetUsers() ([]User, error) {

	rows, err := DB.Query("SELECT * FROM User")
	if err != nil {

	}

	users := make([]User, 0)

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.Name, &u.Login, &u.Password); err != nil {
			return []User{}, err
		}
		users = append(users, u)
	}
	return users, nil
}
