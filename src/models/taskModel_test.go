package models

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/satori/go.uuid"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
	"time"
)

func TestCreateTask(t *testing.T) {

	originalDB := database.DB
	database.DB, mock, err = sqlmock.New()
	defer func() { database.DB = originalDB }()

	until, _ := time.Parse(time.UnixDate, "Mon Jun  15 10:53:39 PST 2018")
	currentTime, _ := time.Parse(time.UnixDate, "Mon Jun  11 10:53:39 PST 2018")

	ID, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")
	UserID, _ := uuid.FromString("00000000-0000-0000-0000-000000000011")

	task := Task{
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

	mock.ExpectExec("INSERT INTO tasks").WithArgs(UserID, "TaskOne",
		until, currentTime, currentTime, "Do smth").WillReturnResult(sqlmock.NewResult(1, 1))

	if _, err := CreateTask(task); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetTask(t *testing.T) {

	originalDB := database.DB
	database.DB, mock, err = sqlmock.New()
	defer func() { database.DB = originalDB }()

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
		AddRow(ID.Bytes(), UserID.Bytes(), "TaskOne", until, currentTime, currentTime, "Do smth")

	mock.ExpectQuery("SELECT (.+) FROM tasks").WithArgs(ID).WillReturnRows(rows)

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

	originalDB := database.DB
	database.DB, mock, err = sqlmock.New()
	defer func() { database.DB = originalDB }()

	ID, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectExec("DELETE FROM tasks").WithArgs(
		ID).WillReturnResult(sqlmock.NewResult(1, 1))

	if err := DeleteTask(ID); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

func TestGetTasks(t *testing.T) {

	originalDB := database.DB
	database.DB, mock, err = sqlmock.New()
	defer func() { database.DB = originalDB }()

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
		AddRow(ID.Bytes(), UserID.Bytes(), "TaskOne", until, currentTime, currentTime,
			"Do smth").AddRow(ID.Bytes(), UserID.Bytes(), "TaskOne", until, currentTime, currentTime, "Do smth")

	mock.ExpectQuery("SELECT (.+) FROM tasks").WillReturnRows(rows)

	result, err := GetTasks()

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
