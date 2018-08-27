package models

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/satori/go.uuid"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
)

var restaurantsMockErr error

func TestAddRestaurantToTrip(t *testing.T) {
	originalDB := database.DB
	database.DB, mock, restaurantsMockErr = sqlmock.New()
	defer func() { database.DB = originalDB }()

	restaurantID, _ := uuid.FromString("00000000-0000-0000-0000-000000000002")
	tripID, _ := uuid.FromString("00000000-0000-0000-0000-000000000003")

	if flightMockErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", flightMockErr)
	}

	mock.ExpectExec("INSERT INTO trips_restaurants").WithArgs(restaurantID, tripID).WillReturnResult(sqlmock.NewResult(1, 1))
	if err := AddToTrip(restaurantID, tripID, Restaurant{}); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
