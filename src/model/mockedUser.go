package model

import "github.com/satori/go.uuid"

//MockedCreateUser is mocked AddUser func
func MockedCreateUser(user User) {
	AddUser = func(user User) (uuid.UUID, error) {
		return user.ID, nil
	}
}

//MockedGetUserByID is mocked GetUserByID func
func MockedGetUserByID(user User, err error) {
	GetUserByID = func(id uuid.UUID) (User, error) {
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
