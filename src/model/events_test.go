package model

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/satori/go.uuid"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
)

var eventMockErr error

func TestAddEventToTrip(t *testing.T) {
	originalDB := database.DB
	database.DB, mock, eventMockErr = sqlmock.New()
	defer func() { database.DB = originalDB }()

	TripID, _ := uuid.FromString("00000000-0000-0000-0000-000000000002")
	EventID, _ := uuid.FromString("00000000-0000-0000-0000-000000000003")

	if eventMockErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", eventMockErr)
	}

	mock.ExpectExec("INSERT INTO trips_events").WithArgs(EventID, TripID).WillReturnResult(sqlmock.NewResult(1, 1))
	if err := AddToTrip(EventID, TripID, Event{}); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
