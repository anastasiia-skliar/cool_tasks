package models

import (
	. "github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/satori/go.uuid"
)

const (
<<<<<<< HEAD
	createUser = "INSERT INTO user (name, login, password) VALUES ($1, $2, $3) RETURNING id"
	getUser    = "SELECT * FROM user WHERE id = $1"
	getUserByID = "SELECT ID, Password FROM Users WHERE Login = $1"
	deleteUser = "DELETE FROM user WHERE id = $1"
	getUsers   = "SELECT * FROM user"
=======
	createUser   = "INSERT INTO users (name, login, password) VALUES ($1, $2, $3) RETURNING id"
	getUser      = "SELECT * FROM users WHERE id = $1"
	deleteUser   = "DELETE FROM users WHERE id = $1"
	getUsers     = "SELECT * FROM users"
>>>>>>> master
)

//User representation in DB
type User struct {
	ID       uuid.UUID
	Name     string
	Login    string
	Password string
}

//CreateUser used for creation user in DB
var CreateUser = func(user User) (uuid.UUID, error) {
	var id uuid.UUID
	err := DB.QueryRow(createUser, user.Name, user.Login, user.Password).Scan(&id)

	return id, err
}

//GetUser used for getting user from DB
var GetUser = func(id uuid.UUID) (User, error) {
	var user User
	err := DB.QueryRow(getUser, id).Scan(&user.ID, &user.Name, &user.Login, &user.Password)

	return user, err
}

<<<<<<< HEAD
//GetUserByLogin used for getting user from DB by Login

func GetUserByLogin(login string) (User, error) {
	var user User
	err:= db.QueryRow(getUserByID, login).Scan(&user.ID, &user.Name, &user.Login, &user.Password)
	return user, err
}

//UpdateUser is used for updating user in DB
func UpdateUser() {

}

=======
>>>>>>> master
//DeleteUser used for deleting user from DB
var DeleteUser = func(id uuid.UUID) error {
	_, err := DB.Exec(deleteUser, id)

	return err
}

//GetUsers used for getting users from DB
var GetUsers = func() ([]User, error) {
	rows, err := DB.Query(getUsers)
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
