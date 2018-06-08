package models

import "database/sql"

var db *sql.DB

type User struct {
	Id int
	Name string
	Login string
	Password string
}

func createUser(user User) error {

	_, err := db.Exec("INSERT INTO user (name, login, password) VALUES ($1, $2, $3)", user.Name, user.Login, user.Password)

	return err
}

func getUser(id int) (*User, error)  {
	user := &User{}
	err := db.QueryRow("SELECT name, login from User WHERE id = $1", user.Id).Scan(user.Name, user.Login)
	if err != nil {
		return user, err
	}

	return user, nil
}

//func updateUser()  {
//
//}

func deleteUser(id int) error {
	_, err := db.Exec("DELETE FROM user WHERE id = $1", id)

	if err != nil{
		return err
	}
	return nil
}

//Метод котрий вертатиме всю інформацію юзерів для Адміа
func getUsersForAdmin() (*[]User, error) {

	rows, err := db.Query("SELECT * FROM user")
	if err != nil {

	}

	users := make([]User, 0)

	for rows.Next() {
		var u User
		if err := rows.Scan(u.Id, u.Name, u.Login, u.Password); err != nil {
			return &[]User{}, err
		}
		users = append(users, u	)
	}
	return &users, nil
}

//Метод котрий вертатиме лише імя та логін юзерів
func getUsers() (*[]User, error) {

	rows, err := db.Query("SELECT name, login FROM user")
	if err != nil {

	}

	users := make([]User, 0)

	for rows.Next() {
		var u User
		if err := rows.Scan(u.Name, u.Login); err != nil {
			return &[]User{}, err
		}
		users = append(users, u	)
	}
	return &users, nil
}
