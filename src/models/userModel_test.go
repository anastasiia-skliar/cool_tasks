package models

import (
	"fmt"
	"github.com/satori/go.uuid"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
)

var mock sqlmock.Sqlmock
var err error

func TestCreateUser(t *testing.T) {

	originalDB := db
	db, mock, err = sqlmock.New()
	defer func() { db = originalDB }()

	UserId, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")

	user := User{
		UserId,
		"John",
		"john",
		"1111",
	}

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"ID"}).AddRow(UserId.Bytes())

	mock.ExpectQuery("INSERT INTO user").WithArgs("John", "john", "1111").WillReturnRows(rows)
	if _, err := CreateUser(user); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetUser(t *testing.T) {

	originalDB := db
	db, mock, err = sqlmock.New()
	defer func() { db = originalDB }()

	UserId, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")

	expected := User{
		ID:       UserId,
		Name:     "John",
		Login:    "john",
		Password: "1111",
	}

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"ID", "Name", "Login", "Password"}).
		AddRow(UserId.Bytes(), "John", "john", "1111")

	mock.ExpectQuery("SELECT (.+) FROM user").WithArgs(UserId).WillReturnRows(rows)

	result, err := GetUser(UserId)

	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	if result != expected {
		t.Error("Expected:", expected, "Was:", result)
	} else {
		fmt.Println(result)
	}
}

func TestDeleteUser(t *testing.T) {

	originalDB := db
	db, mock, err = sqlmock.New()
	defer func() { db = originalDB }()

	id, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectExec("DELETE FROM user WHERE").WithArgs(
		id).WillReturnResult(sqlmock.NewResult(1, 1))

	if err := DeleteUser(id); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetUsers(t *testing.T) {

	originalDB := db
	db, mock, err = sqlmock.New()
	defer func() { db = originalDB }()

	UserId, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")

	var expects = []User{
		{
			ID:       UserId,
			Name:     "John",
			Login:    "john_doe",
			Password: "1111",
		},
		{
			ID:       UserId,
			Name:     "Tom",
			Login:    "hate_jerry",
			Password: "2222",
		},
	}

	rows := sqlmock.NewRows([]string{"ID", "Name", "Login", "Password"}).
		AddRow(UserId.Bytes(), "John", "john_doe", "1111").AddRow(UserId.Bytes(), "Tom", "hate_jerry", "2222")

	mock.ExpectQuery("SELECT (.+) FROM user").WillReturnRows(rows)

	result, err := GetUsers()

	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	for i := 0; i < len(result); i++ {
		if expects[i] != result[i] {
			t.Error("Expected:", expects[i], "Was:", result[i])
		}
	}
}
