package models

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/satori/go.uuid"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
)

var flightMockErr error

type flightsTestCase struct {
	name string
	url  string
}

//func TestGetFlightsByRequest(t *testing.T) {
//	testCase := []flightsTestCase{
//		{"test with ID param",
//			"localhost:8080/v1/flights?id=00000000-0000-0000-0000-000000000001",
//		},
//		{"test with 1 departure_city & 1 arrival_city",
//			"localhost:8080/v1/flights?departure_city=Lviv&arrival_city=Kyiv",
//		},
//		{
//			"test with 2 departure_city & 1 arrival_city",
//			"localhost:8080/v1/flights?departure_city=Lviv&departure_city=Sokal&arrival_city=Kyiv",
//		},
//		{
//			"test with departure_time, departure_date ,arrival_time,arrival_date,price",
//			"localhost:8080/v1/flights?departure_time=Mon Jun  11 10:53:39 PST 2018&departure_date=Mon Jun  11 10:53:39 PST 2018" +
//				"&arrival_time=Mon Jun  12 9:53:39 PST 2018&price=100",
//		},
//		{
//			"test with departure_time, departure_date ,arrival_time,arrival_date, 2 price",
//			"localhost:8080/v1/flights?departure_time=Mon Jun  11 10:53:39 PST 2018&departure_date=Mon Jun  11 10:53:39 PST 2018" +
//				"&arrival_time=Mon Jun  12 9:53:39 PST 2018&price=100&price=200",
//		},
//	}
//
//	originalDB := database.DB
//	database.DB, mock, flightMockErr = sqlmock.New()
//	defer func() { database.DB = originalDB }()
//
//	departureTime, _ := time.Parse(time.UnixDate, "Mon Jun  11 10:53:39 PST 2018")
//	arrivalTime, _ := time.Parse(time.UnixDate, "Mon Jun  12 9:53:39 PST 2018")
//
//	ID, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")
//
//	expected := []Flight{
//		{
//			ID,
//			"Lviv",
//			departureTime,
//			departureTime,
//			"Kyiv",
//			arrivalTime,
//			arrivalTime,
//			100,
//		},
//		{
//			ID,
//			"Sokal",
//			departureTime,
//			departureTime,
//			"Mosty",
//			arrivalTime,
//			arrivalTime,
//			200,
//		},
//	}
//	for _, tc := range testCase {
//		rawUrl, _ := url.Parse(tc.url)
//		params := rawUrl.Query()
//
//		if flightMockErr != nil {
//			t.Fatalf("an error '%s' was not expected when opening a stub database connection", flightMockErr)
//		}
//
//		rows := sqlmock.NewRows([]string{"ID", "departure_city", "departure_time", "departure_date", "arrival_city", "arrival_time", "arrival_date", "price"}).
//			AddRow(ID.Bytes(), "Lviv", departureTime, departureTime, "Kyiv", arrivalTime, arrivalTime, 100).
//			AddRow(ID.Bytes(), "Sokal", departureTime, departureTime, "Mosty", arrivalTime, arrivalTime, 200)
//
//		mock.ExpectQuery("SELECT (.+) FROM flights").WillReturnRows(rows)
//
//		result, err := GetData(params,)
//
//		if err != nil {
//			t.Errorf("error was not expected while updating stats: %s", err)
//		}
//		if err := mock.ExpectationsWereMet(); err != nil {
//			t.Errorf("there were unfulfilled expectations: %s", err)
//		}
//
//		for i := 0; i < len(result); i++ {
//			if expected[i] != result[i] {
//				t.Error("Expected:", expected[i], "Was:", result[i])
//			}
//		}
//
//	}
//
//}

func TestAddFlightToTrip(t *testing.T) {
	originalDB := database.DB
	database.DB, mock, flightMockErr = sqlmock.New()
	defer func() { database.DB = originalDB }()

	flightID, _ := uuid.FromString("00000000-0000-0000-0000-000000000002")
	tripID, _ := uuid.FromString("00000000-0000-0000-0000-000000000003")

	if flightMockErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", flightMockErr)
	}

	mock.ExpectExec("INSERT INTO trips_flights").WithArgs(flightID, tripID).WillReturnResult(sqlmock.NewResult(1, 1))
	if err := AddToTrip(flightID, tripID, Flight{}); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

//func TestGetFlightsByTrip(t *testing.T) {
//	originalDB := database.DB
//	database.DB, mock, flightMockErr = sqlmock.New()
//	defer func() { database.DB = originalDB }()
//
//	departureTime, _ := time.Parse(time.UnixDate, "Mon Jun  11 10:53:39 PST 2018")
//	arrivalTime, _ := time.Parse(time.UnixDate, "Mon Jun  12 9:53:39 PST 2018")
//
//	ID, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")
//
//	expected := Flight{
//
//		ID,
//		"Lviv",
//		departureTime,
//		departureTime,
//		"Kyiv",
//		arrivalTime,
//		arrivalTime,
//		100,
//	}
//
//	if flightMockErr != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", flightMockErr)
//	}
//
//	rows := sqlmock.NewRows([]string{"ID", "departure_city", "departure_time", "departure_date", "arrival_city", "arrival_time", "arrival_date", "price"}).
//		AddRow(ID.Bytes(), "Lviv", departureTime, departureTime, "Kyiv", arrivalTime, arrivalTime, 100)
//
//	mock.ExpectQuery("SELECT (.+) FROM flights INNER JOIN trips_flights ON flights.id=trips_flights.flight_id AND trips_flights.trip_id=\\$1").WithArgs(expected.ID).WillReturnRows(rows)
//
//	result, err := GetFlightsByTrip(ID)
//
//	if err != nil {
//		t.Errorf("error was not expected while updating stats: %s", err)
//	}
//	if err := mock.ExpectationsWereMet(); err != nil {
//		t.Errorf("there were unfulfilled expectations: %s", err)
//	}
//
//	if expected != result[0] {
//		t.Error("Expected:", expected, "Was:", result[0])
//	}
//}
