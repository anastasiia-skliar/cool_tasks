package models

import (
	"github.com/satori/go.uuid"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
	"log"
)

var UUID uuid.UUID

var (
 	errUserMock error
	userMock   sqlmock.Sqlmock
 )

func init()  {
	_, userMock, errUserMock = sqlmock.New()
	if errUserMock != nil {
		log.Fatal(errUserMock)
	}
}

func TestCreateUser(t *testing.T) {

	user := User{
		UUID,
		"John",
		"john",
		"1111",
	}

	if errUserMock != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", errUserMock)
	}

	userMock.ExpectExec("INSERT INTO User").WithArgs(
		"John", "john", "1111").WillReturnResult(sqlmock.NewResult(1, 1))

	if err := CreateUser(user); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := userMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetUser(t *testing.T) {

	id, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")

	expected := User{
		ID: id,
		Name:     "John",
		Login:    "john",
		Password: "1111",
	}

	if errUserMock != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", errUserMock)
	}


	rows := sqlmock.NewRows([]string{"ID", "Name", "Login", "Password"}).
		AddRow( id, "John", "john", "1111")


	userMock.ExpectQuery("SELECT name, login, password FROM User").WithArgs(UUID).WillReturnRows(rows)
	result, err := GetUser(id)

	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
	if err := userMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	if result != expected {
		t.Error("Expected:", expected, "Was:", result)
	}

}

func TestDeleteUser(t *testing.T) {

	id, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")

	if errUserMock != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", errUserMock)
	}

	userMock.ExpectExec("DELETE FROM User WHERE").WithArgs(
		id).WillReturnResult(sqlmock.NewResult(1, 1))

	if err := DeleteUser(id); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
	if err := userMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetUsers(t *testing.T) {
	var expects = []User{
		{
			Name:     "John",
			Login:    "john_doe",
			Password: "1111",
		},
		{
			Name:     "Tom",
			Login:    "hate_jerry",
			Password: "2222",
		},
	}

	rows := sqlmock.NewRows([]string{"Name", "Login", "Password"}).
		AddRow("John", "john_doe", "1111").AddRow("Tom", "hate_jerry", "2222")

	userMock.ExpectQuery("SELECT (.+) FROM User").WillReturnRows(rows)

	result, err := GetUsers()

	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
	if err := userMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	for i := 0; i < len(result); i++ {
		if expects[i] != result[i]{
			t.Error("Expected:", expects[i], "Was:", result[i])
		}
	}
}
