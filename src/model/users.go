package model

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"

	"github.com/satori/go.uuid"
)

const (
	createUser     = "INSERT INTO users (name, login, password, role) VALUES ($1, $2, $3, $4) RETURNING id"
	getUserByID    = "SELECT * FROM users WHERE id = $1"
	usersloc       = "users"
	getUserByLogin = "SELECT * FROM users WHERE login = $1"
	deleteUser     = "DELETE FROM users WHERE id = $1"
	getUsers       = "SELECT * FROM users"
)

//User representation in DB
type User struct {
	ID       uuid.UUID
	Name     string
	Login    string
	Password string
	Role     string
}

//AddUser used for creation user in DB
var AddUser = func(user User) (uuid.UUID, error) {
	var id uuid.UUID
	err := database.DB.QueryRow(createUser, user.Name, user.Login, user.Password, user.Role).Scan(&id)
	return id, err
}

//GetUserByID used for getting user from DB
var GetUserByID = func(id uuid.UUID) (User, error) {
	var user User
	err := database.DB.QueryRow(getUserByID, id).Scan(&user.ID, &user.Name, &user.Login, &user.Password, &user.Role)
	return user, err
}

//GetUserByLogin used for getting user from DB by Login
var GetUserByLogin = func(login string) (User, error) {
	var user User
	err := database.DB.QueryRow(getUserByLogin, login).Scan(&user.ID, &user.Name, &user.Login, &user.Password, &user.Role)
	return user, err
}

//GetUserForLogin used for getting user from DB by Login and Password
var GetUserForLogin = func(login string, password string) (User, error) {
	var user User
	params := []string{"login", "password"}
	values := make(map[string][]string)
	values["login"] = []string{login}
	values["password"] = []string{password}
	req, sqlParams, err := SQLGenerator(usersloc, params, nil, values)
	if err != nil {
		return User{}, err
	}
	err = database.DB.QueryRow(req, sqlParams...).Scan(&user.ID, &user.Name, &user.Login, &user.Password, &user.Role)
	return user, err
}

//DeleteUser used for deleting user from DB
var DeleteUser = func(id uuid.UUID) error {
	_, err := database.DB.Exec(deleteUser, id)
	return err
}

//GetUsers used for getting users from DB
var GetUsers = func() ([]User, error) {
	rows, err := database.DB.Query(getUsers)
	if err != nil {
		return nil, err
	}

	users := make([]User, 0)

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Login, &u.Password, &u.Role); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}
