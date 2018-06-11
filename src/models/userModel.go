package models

import (
	"github.com/satori/go.uuid"
	"database/sql"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"log"
)

type User struct {
	ID       uuid.UUID
	Name     string
	Login    string
	Password string
}

//will be deleteted!
var (
	db *sql.DB
	err error
)

func init()  {
	db, _, err = sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
}

func CreateUser(user User) error {
	_, err := db.Exec("INSERT INTO User (name, login, password) VALUES ($1, $1, $3)", user.Name, user.Login, user.Password)

	if err != nil {
		return err
	}
	return nil
}

func GetUser(id uuid.UUID) (User, error) {
	user := User{}
	err := db.QueryRow("SELECT name, login, password FROM User WHERE id = $1", id).Scan(&user.Name, &user.Login, &user.Password)

	if err != nil {
		return user, err
	}
	return user, nil
}

func UpdateUser() {

}

func DeleteUser(id uuid.UUID) error {
	_, err := db.Exec("DELETE FROM User WHERE id = $1", id)

	if err != nil {
		return err
	}
	return nil
}

func GetUsers() ([]User, error) {

	rows, err := db.Query("SELECT * FROM User")
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