package models_test

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/Nastya-Kruglikova/cool_tasks/src/models"
	"github.com/satori/go.uuid"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
)

type TripsTestCase struct {
	name                   string
	mockedGetTripsByTripID models.Trip
}

func TestCreateTrip(t *testing.T) {
	originalDB := database.DB
	database.DB, mock, mockErr = sqlmock.New()
	defer func() { database.DB = originalDB }()

	UserID, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")
	TripID, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")

	trip := models.Trip{
		TripID: TripID,
		UserID: UserID,
	}

	if mockErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", mockErr)
	}

	rows := sqlmock.NewRows([]string{"ID"}).AddRow(UserID.Bytes())

	mock.ExpectQuery("INSERT INTO trips").WithArgs(UserID).WillReturnRows(rows)
	if _, err := models.CreateTrip(trip); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetTripsByUserID(t *testing.T) {
	originalDB := database.DB
	database.DB, mock, mockErr = sqlmock.New()
	defer func() { database.DB = originalDB }()

	ID, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")
	TripID, _ := uuid.FromString("00000000-0000-0000-0000-000000000002")

	expects := []uuid.UUID{TripID}

	rows := sqlmock.NewRows([]string{"TripID"}).
		AddRow(expects[0].Bytes())

	mock.ExpectQuery("SELECT trips.trip_id FROM trips WHERE trips.user_id").WithArgs(ID).WillReturnRows(rows)

	_, err := models.GetTripIDByUserID(ID)

	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
