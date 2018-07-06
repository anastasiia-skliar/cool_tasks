package models

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/satori/go.uuid"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
	"time"
)

var (
	mock    sqlmock.Sqlmock
	mockErr error
)

func TestGetTrains(t *testing.T) {

	originalDB := database.DB
	database.DB, mock, mockErr = sqlmock.New()
	defer func() { database.DB = originalDB }()

	ID, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")

	departureTime, _ := time.Parse("15:04:05", "12:00:00")
	departureDate, _ := time.Parse("2006-01-02", "2018-07-21")
	arrivalTime, _ := time.Parse("15:04:05", "15:00:00")
	arrivalDate, _ := time.Parse("2006-01-02", "2018-07-21")

	expects := []Train{
		{
			ID,
			departureTime,
			departureDate,
			arrivalTime,
			arrivalDate,
			"Lviv",
			"Kyiv",
			"el",
			"coupe",
			"200uah",
		},
		{
			ID,
			departureTime,
			departureDate,
			arrivalTime,
			arrivalDate,
			"Lviv",
			"Kharkiv",
			"el",
			"coupe",
			"250uah",
		},
	}

	if mockErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", mockErr)
	}

	rows := sqlmock.NewRows([]string{"ID", "departureTime", "departureDate", "arrivalTime", "arrivalDate",
		"departure_city", "arrival_city", "train_type", "car_type", "price"}).
		AddRow(ID.Bytes(), departureTime, departureDate, arrivalTime, arrivalDate,
			"Lviv", "Kyiv", "el", "coupe", "200uah").AddRow(ID.Bytes(), departureTime, departureDate, arrivalTime, arrivalDate,
		"Lviv", "Kharkiv", "el", "coupe", "250uah")

	mock.ExpectQuery("SELECT (.+) FROM trains").WillReturnRows(rows)

	result, err := GetTrains("SELECT (.+) FROM trains")

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

func TestSaveToTrip(t *testing.T) {

	originalDB := database.DB
	database.DB, mock, mockErr = sqlmock.New()
	defer func() { database.DB = originalDB }()

	trainID, _ := uuid.FromString("00000000-0000-0000-0000-000000000002")
	tripID, _ := uuid.FromString("00000000-0000-0000-0000-000000000003")

	if mockErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", mockErr)
	}

	mock.ExpectExec("INSERT INTO trips_trains").WithArgs(trainID, tripID).WillReturnResult(sqlmock.NewResult(1, 1))
	if err := SaveTrain(trainID, tripID); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetFromTrip(t *testing.T) {
	originalDB := database.DB
	database.DB, mock, mockErr = sqlmock.New()
	defer func() { database.DB = originalDB }()

	if mockErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", mockErr)
	}

	departureTime, _ := time.Parse("15:04:05", "12:00:00")
	departureDate, _ := time.Parse("2006-01-02", "2018-07-21")
	arrivalTime, _ := time.Parse("15:04:05", "15:00:00")
	arrivalDate, _ := time.Parse("2006-01-02", "2018-07-21")

	ID, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")

	expects := []Train{
		{
			ID,
			departureTime,
			departureDate,
			arrivalTime,
			arrivalDate,
			"Lviv",
			"Kyiv",
			"el",
			"coupe",
			"200uah",
		},
		{
			ID,
			departureTime,
			departureDate,
			arrivalTime,
			arrivalDate,
			"Lviv",
			"Kharkiv",
			"el",
			"coupe",
			"250uah",
		},
	}

	rows := sqlmock.NewRows([]string{"ID", "departure_time", "departure_date", "arrival_time", "arrival_date",
		"departure_city", "arrival_city", "train_type", "car_type", "price"}).
		AddRow(ID.Bytes(), departureTime, departureDate, arrivalTime, arrivalDate,
			"Lviv", "Kyiv", "el", "coupe", "200uah")

	mock.ExpectQuery("SELECT (.+) FROM trains INNER JOIN trips_trains ON trips_trains.trains_id = trains.id AND trips_trains.trips").WithArgs(ID).WillReturnRows(rows)

	result, err := GetFromTrip(ID)

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
