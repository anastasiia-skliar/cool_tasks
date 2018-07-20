package models_test

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/Nastya-Kruglikova/cool_tasks/src/models"
	"github.com/satori/go.uuid"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
	"time"
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

func TestGetTripsByTripID(t *testing.T) {

	ID, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")
	TripID, _ := uuid.FromString("00000000-0000-0000-0000-000000000002")
	UserID, _ := uuid.FromString("00000000-0000-0000-0000-000000000003")

	departureTime, _ := time.Parse("15:04:05", "12:00:00")
	departureDate, _ := time.Parse("2006-01-02", "2018-07-21")
	arrivalTime, _ := time.Parse("15:04:05", "15:00:00")
	arrivalDate, _ := time.Parse("2006-01-02", "2018-07-21")
	testTime, _ := time.Parse("15:04:05", "12:00:00")
	MuseumId, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")

	expects := models.Trip{
		TripID:  TripID,
		UserID:  UserID,
		Events:  []models.Event{},
		Flights: []models.Flight{},
		Museums: []models.Museum{
			{
				ID:         MuseumId,
				Name:       "Ermitage",
				Location:   "Peterburg",
				Price:      1111,
				OpenedAt:   testTime,
				ClosedAt:   testTime,
				MuseumType: "Gallery",
				Info:       "Cool",
			},
		},
		Trains: []models.Train{
			{
				ID:            ID,
				DepartureTime: departureTime,
				DepartureDate: departureDate,
				ArrivalTime:   arrivalTime,
				ArrivalDate:   arrivalDate,
				DepartureCity: "Lviv",
				ArrivalCity:   "Kyiv",
				TrainType:     "el",
				CarType:       "coupe",
				Price:         "200uah",
			},
		},
	}

	tests := TripsTestCase{
		name: "Get_Trips_OK",
		mockedGetTripsByTripID: models.Trip{},
	}

	t.Run(tests.name, func(t *testing.T) {
		models.MockedGetTripsByTripID(expects)
	})
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
