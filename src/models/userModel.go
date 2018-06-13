package models

import (
	"database/sql"
	"github.com/satori/go.uuid"
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
)

const (
	createUser = "INSERT INTO user (name, login, password) VALUES ($1, $2, $3) RETURNING id"
	getUser    = "SELECT * FROM user WHERE id = $1"
	deleteUser = "DELETE FROM user WHERE id = $1"
	getUsers   = "SELECT * FROM user"
)

//User representation in DB
type User struct {
	ID       uuid.UUID
	Name     string
	Login    string
	Password string
}

var PostgeSQLInfo = database.PostgreSQLInfo{
	"127.0.0.1",
5432,
"cool_tasks",
"postgres",
"12345678"}
var Info = database.Info{PostgreSQL:PostgeSQLInfo}

//will be deleteted!
var db *sql.DB
var err error

func init()  {
	db, err = database.SetupPostgres(Info)
}

//CreateUser used for creation user in DB
func CreateUser(user User) (uuid.UUID, error) {
	var id uuid.UUID
	err := db.QueryRow(createUser, user.Name, user.Login, user.Password).Scan(&id)

	return id, err
}

//GetUser used for getting user from DB
func GetUser(id uuid.UUID) (User, error) {
	var user User
	err := db.QueryRow(getUser, id).Scan(&user.ID, &user.Name, &user.Login, &user.Password)

	return user, err
}

//UpdateUser is used for updating user in DB
func UpdateUser() {

}

//DeleteUser used for deleting user from DB
func DeleteUser(id uuid.UUID) error {
	_, err := db.Exec(deleteUser, id)

	return err
}

//GetUsers used for getting users from DB
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
