package models

import "github.com/satori/go.uuid"

func MockedCreateUser(user User) {
	CreateUser = func(user User) (uuid.UUID, error) {
		return user.ID, nil
	}
}

func MockedGetUser(user User, err error) {
	GetUser = func(id uuid.UUID) (User, error) {
		return user, err
	}
}

func MockedDeleteUser(id uuid.UUID, err error) {
	DeleteUser = func(id uuid.UUID) error {
		return err
	}
}

func MockedGetUsers(user []User, err error) {
	GetUsers = func() ([]User, error) {
		return user, err
	}
}
