package models

import (
	"testing"
	"time"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"github.com/satori/go.uuid"
)

func TestCreateTask(t *testing.T) {

	originalDB := db
	db, mock, err = sqlmock.New()
	defer func() { db = originalDB }()

	until, _ := time.Parse(time.UnixDate, "Mon Jun  15 10:53:39 PST 2018")
	currentTime, _ := time.Parse(time.UnixDate, "Mon Jun  11 10:53:39 PST 2018")

	ID, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")
	UserID, _ := uuid.FromString("00000000-0000-0000-0000-000000000011")

	task :=Task{
		ID,
		UserID,
		"TaskOne",
		until,
		currentTime,
		currentTime,
		"Do smth",
	}

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectExec("INSERT INTO task").WithArgs(UserID, "TaskOne",
		until, currentTime, currentTime, "Do smth").WillReturnResult(sqlmock.NewResult(1, 1))

	if err := CreateTask(task); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetTask(t *testing.T) {

	originalDB := db
	db, mock, err = sqlmock.New()
	defer func() { db = originalDB }()

	ID, _ := uuid.FromString("00000000-0000-0000-0000-00000000001")
	UserID, _ := uuid.FromString("00000000-0000-0000-0000-000000000011")

	until, _ := time.Parse(time.UnixDate, "Mon Jun  15 10:53:39 PST 2018")
	currentTime, _ := time.Parse(time.UnixDate, "Mon Jun  11 10:53:39 PST 2018")

	expected := Task{
		ID:        ID,
		UserID:    UserID,
		Name:      "TaskOne",
		Time:      until,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
		Desc:      "Do smth",
	}

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"ID", "UserID", "Name", "Time", "CreatedAt", "UpdatedAt", "Desc"}).
		AddRow(ID, UserID, "TaskOne", until, currentTime, currentTime, "Do smth")

	mock.ExpectQuery("SELECT (.+) FROM task").WithArgs(ID).WillReturnRows(rows)

	result, err := GetTask(ID)

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

func TestDeleteTask(t *testing.T) {

	ID, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")

	originalDB := db
	db, mock, err = sqlmock.New()
	defer func() { db = originalDB }()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectExec("DELETE FROM task").WithArgs(
		ID).WillReturnResult(sqlmock.NewResult(1, 1))

	if err := DeleteTask(ID); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

func TestGetTasks(t *testing.T) {

	originalDB := db
	db, mock, err = sqlmock.New()
	defer func() { db = originalDB }()

	until, _ := time.Parse(time.UnixDate, "Mon Jun  15 10:53:39 PST 2018")
	currentTime, _ := time.Parse(time.UnixDate, "Mon Jun  11 10:53:39 PST 2018")

	ID, _ := uuid.FromString("00000000-0000-0000-0000-00000000001")
	UserID, _ := uuid.FromString("00000000-0000-0000-0000-000000000011")

	var expects = []Task{
		{
			ID,
			UserID,
			"TaskOne",
			until,
			currentTime,
			currentTime,
			"Do smth",
		},
		{
			ID,
			UserID,
			"TaskOne",
			until,
			currentTime,
			currentTime,
			"Do smth",
		},
	}

	rows := sqlmock.NewRows([]string{"ID", "UserID", "Name", "Time", "CreatedAt", "UpdatedAt", "Desc"}).
		AddRow(ID, UserID, "TaskOne", until, currentTime, currentTime,
			"Do smth").AddRow(ID, UserID, "TaskOne", until, currentTime, currentTime, "Do smth")

	mock.ExpectQuery("SELECT (.+) FROM task").WillReturnRows(rows)

	result, err := GetTasks()

	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	for i := 0; i < len(result); i++ {
		if expects[i] != result[i]{
			t.Error("Expected:", expects[i], "Was:", result[i])
		}
	}
}