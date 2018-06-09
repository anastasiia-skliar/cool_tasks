package models

import (
	"github.com/satori/go.uuid"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
)

var UUID uuid.UUID

func TestCreateUser(t *testing.T) {

	user := User{
		UUID,
		"John",
		"john",
		"1111",
	}

	if Err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", Err)
	}

	Mock.ExpectExec("INSERT INTO User").WithArgs(
		"John", "john", "1111").WillReturnResult(sqlmock.NewResult(1, 1))

	if err := CreateUser(user); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := Mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetUser(t *testing.T) {

	expected := User{
		Name:     "John",
		Login:    "john",
		Password: "1111",
	}

	if Err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", Err)
	}

	rows := sqlmock.NewRows([]string{"Name", "Login", "Password"}).
		AddRow("John", "john", "1111")
	//rows := sqlmock.NewRows([]string{"Name", "Login", "Password"}).
	//	AddRow(UUID, "John", "john", "1111")

	Mock.ExpectQuery("^SELECT (.+) FROM User WHERE").WithArgs(UUID).WillReturnRows(rows)

	result, err := GetUser(UUID)

	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
	if err := Mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	if *result != expected {
		t.Error("Expected:", expected, "Was:", *result)
	}

}

func TestDeleteUser(t *testing.T) {

	if Err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", Err)
	}

	Mock.ExpectExec("DELETE FROM User").WithArgs(
		UUID).WillReturnResult(sqlmock.NewResult(1, 1))

	if err := DeleteUser(UUID); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
	if err := Mock.ExpectationsWereMet(); err != nil {
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

	Mock.ExpectQuery("SELECT (.+) FROM User").WillReturnRows(rows)

	result, err := GetUsers()

	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
	if err := Mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	for i := 0; i < len(result); i++ {
		if expects[i] != result[i]{
			t.Error("Expected:", expects[i], "Was:", result[i])
		}
	}
}
