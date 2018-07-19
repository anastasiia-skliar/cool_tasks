package models

import "github.com/satori/go.uuid"

//MockedCreateUser is mocked CreateUser func
func MockedCreateUser(user User) {
	CreateUser = func(user User) (uuid.UUID, error) {
		return user.ID, nil
	}
}

//MockedGetUser is mocked GetUser func
func MockedGetUser(user User, err error) {
	GetUser = func(id uuid.UUID) (User, error) {
		return user, err
	}
}

//MockedDeleteUser is mocked DeleteUser func
func MockedDeleteUser(id uuid.UUID, err error) {
	DeleteUser = func(id uuid.UUID) error {
		return err
	}
}

//MockedGetUsers is mocked GetUsers func
func MockedGetUsers(user []User, err error) {
	GetUsers = func() ([]User, error) {
		return user, err
	}
}
