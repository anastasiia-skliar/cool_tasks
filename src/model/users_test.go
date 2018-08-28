package model

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/satori/go.uuid"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
)

var mock sqlmock.Sqlmock
var userMockErr error

func TestCreateUser(t *testing.T) {

	originalDB := database.DB
	database.DB, mock, userMockErr = sqlmock.New()
	defer func() { database.DB = originalDB }()

	UserID, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")

	user := User{
		UserID,
		"John",
		"john",
		"1111",
		"Admin",
	}

	if userMockErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", userMockErr)
	}

	rows := sqlmock.NewRows([]string{"ID"}).AddRow(UserID.Bytes())

	mock.ExpectQuery("INSERT INTO users").WithArgs("John", "john", "1111", "Admin").WillReturnRows(rows)
	if _, err := AddUser(user); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetUserByID(t *testing.T) {

	originalDB := database.DB
	database.DB, mock, userMockErr = sqlmock.New()
	defer func() { database.DB = originalDB }()

	UserID, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")

	expected := User{
		ID:       UserID,
		Name:     "John",
		Login:    "john",
		Password: "1111",
		Role:     "Admin",
	}

	if userMockErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", userMockErr)
	}

	rows := sqlmock.NewRows([]string{"ID", "Name", "Login", "Password", "Role"}).
		AddRow(UserID.Bytes(), "John", "john", "1111", "Admin")

	mock.ExpectQuery("SELECT (.+) FROM users").WithArgs(UserID).WillReturnRows(rows)

	result, err := GetUserByID(UserID)

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
	database.DB, mock, userMockErr = sqlmock.New()
	defer func() { database.DB = originalDB }()

	UserID, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")

	expected := User{
		ID:       UserID,
		Name:     "John",
		Login:    "john",
		Password: "1111",
		Role:     "Admin",
	}

	if userMockErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", userMockErr)
	}

	rows := sqlmock.NewRows([]string{"ID", "Name", "Login", "Password", "Role"}).
		AddRow(UserID.Bytes(), "John", "john", "1111", "Admin")

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
	database.DB, mock, userMockErr = sqlmock.New()
	defer func() { database.DB = originalDB }()

	id, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")

	if userMockErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", userMockErr)
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
	database.DB, mock, userMockErr = sqlmock.New()
	defer func() { database.DB = originalDB }()

	UserID, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")

	var expects = []User{
		{
			ID:       UserID,
			Name:     "John",
			Login:    "john_doe",
			Password: "****",
			Role:     "Admin",
		},
		{
			ID:       UserID,
			Name:     "Tom",
			Login:    "hate_jerry",
			Password: "****",
			Role:     "Admin",
		},
	}

	if userMockErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", userMockErr)
	}

	rows := sqlmock.NewRows([]string{"ID", "Name", "Login", "Password", "Role"}).
		AddRow(UserID.Bytes(), "John", "john_doe", "****", "Admin").AddRow(UserID.Bytes(), "Tom", "hate_jerry", "****", "Admin")

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
