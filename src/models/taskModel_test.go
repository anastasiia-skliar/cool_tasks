package models

import (
	"testing"
	"time"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"github.com/satori/go.uuid"
	"database/sql"
	"log"
)

var UUIDu uuid.UUID

var (
	dbTaskMock          *sql.DB
	errTaskMock 		  error
	taskMock    sqlmock.Sqlmock
)

func init()  {
	db, userMock, err = sqlmock.New()
	if errUserMock != nil {
		log.Fatal(errUserMock)
	}
}

func TestCreateTask(t *testing.T) {

	until, err := time.Parse(time.UnixDate, "Mon Jun  15 10:53:39 PST 2018")
	if err != nil {
		panic(err)
	}

	currentTime, err := time.Parse(time.UnixDate, "Mon Jun  11 10:53:39 PST 2018")
	if err != nil {
		panic(err)
	}

	task :=Task{
		UUID,
		UUIDu,
		"TaskOne",
		until,
		currentTime,
		currentTime,
		"Do smth",
	}

	if errTaskMock != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", errUserMock)
	}

	userMock.ExpectExec("INSERT INTO task").WithArgs(UUIDu, "TaskOne",
		until, currentTime, currentTime, "Do smth").WillReturnResult(sqlmock.NewResult(1, 1))

	if err := CreateTask(task); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := taskMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetTask(t *testing.T) {

	until, err := time.Parse(time.UnixDate, "Mon Jun  15 10:53:39 PST 2018")
	if err != nil {
		panic(err)
	}

	currentTime, err := time.Parse(time.UnixDate, "Mon Jun  11 10:53:39 PST 2018")
	if err != nil {
		panic(err)
	}

	expected := Task{
		User_ID: 		  UUIDu,
		Name: 		  "TaskOne",
		Time: 			  until,
		Created_at: currentTime,
		Updated_at: currentTime,
		Desc: 		  "Do smth",
	}


	if errTaskMock != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", errTaskMock)
	}

	rows := sqlmock.NewRows([]string{"User_ID", "Name", "Time", "Created_at", "Updated_at", "Desc"}).
		AddRow(UUIDu, "TaskOne", until, currentTime, currentTime, "Do smth")

	taskMock.ExpectQuery("SELECT user_id, name, time, created_at, updated_at, desc FROM task").WithArgs(UUID).WillReturnRows(rows)

	result, err := GetTask(UUID)

	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
	if err := taskMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	if result != expected {
		t.Error("Expected:", expected, "Was:", result)
	}

}

func TestDeleteTask(t *testing.T) {

	if errTaskMock != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", errTaskMock)
	}

	taskMock.ExpectExec("DELETE FROM task").WithArgs(
		UUID).WillReturnResult(sqlmock.NewResult(1, 1))

	if err := DeleteTask(UUID); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
	if err := taskMock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

func TestGetTasks(t *testing.T) {


}