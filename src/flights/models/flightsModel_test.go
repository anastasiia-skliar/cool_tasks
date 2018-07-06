package models

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/satori/go.uuid"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
	"time"
)

var mock sqlmock.Sqlmock
var err error

func TestGetFlights(t *testing.T) {

	originalDB := database.DB
	database.DB, mock, err = sqlmock.New()
	defer func() { database.DB = originalDB }()

	departureTime, _ := time.Parse(time.UnixDate, "Mon Jun  11 10:53:39 PST 2018")
	arrivalTime, _ := time.Parse(time.UnixDate, "Mon Jun  12 9:53:39 PST 2018")

	ID, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")

	expected := []Flights{
		{
			ID,
			"Lviv",
			departureTime,
			departureTime,
			"Kyiv",
			arrivalTime,
			arrivalTime,
			100,
		},
		{
			ID,
			"Sokal",
			departureTime,
			departureTime,
			"Mosty",
			arrivalTime,
			arrivalTime,
			200,
		},
	}

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"ID", "departure_city", "departure_time", "departure_date", "arrival_city", "arrival_time", "arrival_date", "price"}).

		AddRow(ID.Bytes(), "Lviv", departureTime, departureTime, "Kyiv", arrivalTime, arrivalTime, 100).
		AddRow(ID.Bytes(), "Sokal", departureTime, departureTime, "Mosty", arrivalTime, arrivalTime, 200)

	mock.ExpectQuery("SELECT (.+) FROM flights").WillReturnRows(rows)

	result, err := GetFlights()

	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	for i := 0; i < len(result); i++ {
		if expected[i] != result[i] {
			t.Error("Expected:", expected[i], "Was:", result[i])
		}
	}

}

func TestGetByCity(t *testing.T) {

	originalDB := database.DB
	database.DB, mock, err = sqlmock.New()
	defer func() { database.DB = originalDB }()

	departureTime, _ := time.Parse(time.UnixDate, "Mon Jun  11 10:53:39 PST 2018")
	arrivalTime, _ := time.Parse(time.UnixDate, "Mon Jun  12 9:53:39 PST 2018")

	ID, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")

	expected := Flights{

		ID,
		"Lviv",
		departureTime,
		departureTime,
		"Kyiv",
		arrivalTime,
		arrivalTime,
		100,
	}

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"ID", "departure_city", "departure_time", "departure_date", "arrival_city", "arrival_time", "arrival_date", "price"}).

		AddRow(ID.Bytes(), "Lviv", departureTime, departureTime, "Kyiv", arrivalTime, arrivalTime, 100)

	mock.ExpectQuery("SELECT (.+) FROM flights WHERE departure_city=\\$1 AND arrival_city=\\$2").WithArgs(expected.departureCity, expected.arrivalCity).WillReturnRows(rows)

	result, err := GetByCity(expected.departureCity, expected.arrivalCity)

	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	if expected != result[0] {
		t.Error("Expected:", expected, "Was:", result[0])
	}

}

func TestGetByDepartureTime(t *testing.T) {

	originalDB := database.DB
	database.DB, mock, err = sqlmock.New()
	defer func() { database.DB = originalDB }()

	departureTime, _ := time.Parse(time.UnixDate, "Mon Jun  11 10:53:39 PST 2018")
	departureTimeFrom, _ := time.Parse(time.UnixDate, "Mon Jun  10 10:53:39 PST 2018")
	departureTimeTo, _ := time.Parse(time.UnixDate, "Mon Jun  12 10:53:39 PST 2018")
	arrivalTime, _ := time.Parse(time.UnixDate, "Mon Jun  12 9:53:39 PST 2018")

	ID, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")

	expected := Flights{

		ID,
		"Lviv",
		departureTime,
		departureTime,
		"Kyiv",
		arrivalTime,
		arrivalTime,
		100,
	}

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"ID", "departure_city", "departure_time", "departure_date", "arrival_city", "arrival_time", "arrival_date", "price"}).

		AddRow(ID.Bytes(), "Lviv", departureTime, departureTime, "Kyiv", arrivalTime, arrivalTime, 100)

	mock.ExpectQuery("SELECT (.+) FROM flights WHERE departure_time BETWEEN \\$1 AND \\$2").WithArgs(departureTimeFrom, departureTimeTo).WillReturnRows(rows)

	result, err := GetByDepartureTime(departureTimeFrom, departureTimeTo)

	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	if expected != result[0] {
		t.Error("Expected:", expected, "Was:", result[0])
	}

}

func TestGetByArrivalTime(t *testing.T) {

	originalDB := database.DB
	database.DB, mock, err = sqlmock.New()
	defer func() { database.DB = originalDB }()

	departureTime, _ := time.Parse(time.UnixDate, "Mon Jun  11 10:53:39 PST 2018")
	arrivalTimeFrom, _ := time.Parse(time.UnixDate, "Mon Jun  11 10:53:39 PST 2018")
	arrivalTimeTo, _ := time.Parse(time.UnixDate, "Mon Jun  13 10:53:39 PST 2018")
	arrivalTime, _ := time.Parse(time.UnixDate, "Mon Jun  12 9:53:39 PST 2018")

	ID, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")

	expected := Flights{

		ID,
		"Lviv",
		departureTime,
		departureTime,
		"Kyiv",
		arrivalTime,
		arrivalTime,
		100,
	}

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"ID", "departure_city", "departure_time", "departure_date", "arrival_city", "arrival_time", "arrival_date", "price"}).

		AddRow(ID.Bytes(), "Lviv", departureTime, departureTime, "Kyiv", arrivalTime, arrivalTime, 100)

	mock.ExpectQuery("SELECT (.+) FROM flights WHERE arrival_time BETWEEN \\$1 AND \\$2").WithArgs(arrivalTimeFrom, arrivalTimeTo).WillReturnRows(rows)

	result, err := GetByArrivalTime(arrivalTimeFrom, arrivalTimeTo)

	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	if expected != result[0] {
		t.Error("Expected:", expected, "Was:", result[0])
	}

}

func TestGetByPrice(t *testing.T) {

	originalDB := database.DB
	database.DB, mock, err = sqlmock.New()
	defer func() { database.DB = originalDB }()

	departureTime, _ := time.Parse(time.UnixDate, "Mon Jun  11 10:53:39 PST 2018")
	arrivalTime, _ := time.Parse(time.UnixDate, "Mon Jun  12 9:53:39 PST 2018")

	ID, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")

	expected := Flights{

		ID,
		"Lviv",
		departureTime,
		departureTime,
		"Kyiv",
		arrivalTime,
		arrivalTime,
		100,
	}

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"ID", "departure_city", "departure_time", "departure_date", "arrival_city", "arrival_time", "arrival_date", "price"}).

		AddRow(ID.Bytes(), "Lviv", departureTime, departureTime, "Kyiv", arrivalTime, arrivalTime, 100)

	mock.ExpectQuery("SELECT (.+) FROM flights WHERE price BETWEEN \\$1 AND \\$2").WithArgs(50, 150).WillReturnRows(rows)

	result, err := GetByPrice(50, 150)

	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	if expected != result[0] {
		t.Error("Expected:", expected, "Was:", result[0])
	}

}

func TestToTrip(t *testing.T) {
	originalDB := database.DB
	database.DB, mock, err = sqlmock.New()
	defer func() { database.DB = originalDB }()

	flightID, _ := uuid.FromString("00000000-0000-0000-0000-000000000002")
	tripID, _ := uuid.FromString("00000000-0000-0000-0000-000000000003")

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectExec("INSERT INTO trips_flights").WithArgs(flightID, tripID).WillReturnResult(sqlmock.NewResult(1, 1))
	if err := AddToTrip(flightID, tripID); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetByDate(t *testing.T) {

	originalDB := database.DB
	database.DB, mock, err = sqlmock.New()
	defer func() { database.DB = originalDB }()

	departureTime, _ := time.Parse(time.UnixDate, "Mon Jun  11 10:53:39 PST 2018")
	dateFrom, _ := time.Parse(time.UnixDate, "Mon Jun  10 10:53:39 PST 2018")
	dateTo, _ := time.Parse(time.UnixDate, "Mon Jun  13 10:53:39 PST 2018")
	arrivalTime, _ := time.Parse(time.UnixDate, "Mon Jun  12 9:53:39 PST 2018")

	ID, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")

	expected := Flights{

		ID,
		"Lviv",
		departureTime,
		departureTime,
		"Kyiv",
		arrivalTime,
		arrivalTime,
		100,
	}

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"ID", "departure_city", "departure_time", "departure_date", "arrival_city", "arrival_time", "arrival_date", "price"}).

		AddRow(ID.Bytes(), "Lviv", departureTime, departureTime, "Kyiv", arrivalTime, arrivalTime, 100)

	mock.ExpectQuery("SELECT (.+) FROM flights WHERE departure_date = \\$1 AND arrival_date = \\$2").WithArgs(dateFrom, dateTo).WillReturnRows(rows)

	result, err := GetByDate(dateFrom, dateTo)

	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

	if expected != result[0] {
		t.Error("Expected:", expected, "Was:", result[0])
	}

}