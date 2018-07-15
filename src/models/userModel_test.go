package models

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/satori/go.uuid"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
)

var mock sqlmock.Sqlmock
var mockErr error

func TestCreateUser(t *testing.T) {

	originalDB := database.DB
	database.DB, mock, mockErr = sqlmock.New()
	defer func() { database.DB = originalDB }()

	UserID, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")

	user := User{
		UserID,
		"John",
		"john",
		"1111",
	}

	if mockErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", mockErr)
	}

	rows := sqlmock.NewRows([]string{"ID"}).AddRow(UserID.Bytes())

	mock.ExpectQuery("INSERT INTO users").WithArgs("John", "john", "1111").WillReturnRows(rows)
	if _, err := CreateUser(user); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetUser(t *testing.T) {

	originalDB := database.DB
	database.DB, mock, mockErr = sqlmock.New()
	defer func() { database.DB = originalDB }()

	UserID, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")

	expected := User{
		ID:       UserID,
		Name:     "John",
		Login:    "john",
		Password: "****",
	}

	if mockErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", mockErr)
	}

	rows := sqlmock.NewRows([]string{"ID", "Name", "Login", "Password"}).
		AddRow(UserID.Bytes(), "John", "john", "****")

	mock.ExpectQuery("SELECT (.+) FROM users").WithArgs(UserID).WillReturnRows(rows)

	result, err := GetUser(UserID)

	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	if result != expected {
		t.Error("Expected:", expected, "Was:", result)
	}
}

func TestGetUserByLogin(t *testing.T) {

	originalDB := database.DB
	database.DB, mock, mockErr = sqlmock.New()
	defer func() { database.DB = originalDB }()

	UserID, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")

	expected := User{
		ID:       UserID,
		Name:     "John",
		Login:    "john",
		Password: "1111",
	}

	if mockErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", mockErr)
	}

	rows := sqlmock.NewRows([]string{"ID", "Name", "Login", "Password"}).
		AddRow(UserID.Bytes(), "John", "john", "1111")

	mock.ExpectQuery("SELECT (.+) FROM users WHERE login").WithArgs(expected.Login).WillReturnRows(rows)

	result, err := GetUserByLogin(expected.Login)

	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	if result != expected {
		t.Error("Expected:", expected, "Was:", result)
	}
}

func TestDeleteUser(t *testing.T) {

	originalDB := database.DB
	database.DB, mock, mockErr = sqlmock.New()
	defer func() { database.DB = originalDB }()

	id, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")

	if mockErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", mockErr)
	}

	mock.ExpectExec("DELETE FROM users WHERE").WithArgs(
		id).WillReturnResult(sqlmock.NewResult(1, 1))

	if err := DeleteUser(id); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetUsers(t *testing.T) {

	originalDB := database.DB
	database.DB, mock, mockErr = sqlmock.New()
	defer func() { database.DB = originalDB }()

	UserID, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")

	var expects = []User{
		{
			ID:       UserID,
			Name:     "John",
			Login:    "john_doe",
			Password: "1111",
		},
		{
			ID:       UserID,
			Name:     "Tom",
			Login:    "hate_jerry",
			Password: "2222",
		},
	}

	rows := sqlmock.NewRows([]string{"ID", "Name", "Login", "Password"}).
		AddRow(UserID.Bytes(), "John", "john_doe", "1111").AddRow(UserID.Bytes(), "Tom", "hate_jerry", "2222")

	mock.ExpectQuery("SELECT (.+) FROM users").WillReturnRows(rows)

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
