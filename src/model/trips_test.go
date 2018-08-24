package model_test

import (
	"fmt"
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/Nastya-Kruglikova/cool_tasks/src/model"
	"github.com/satori/go.uuid"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
)

type TripsTestCase struct {
	name                   string
	mockedGetTripsByTripID model.Trip
	mockedTripError        error
	expectedTripId         uuid.UUID
	mock                   func()
}

func TestCreateTrip(t *testing.T) {
	originalDB := database.DB
	database.DB, mock, mockErr = sqlmock.New()
	defer func() { database.DB = originalDB }()

	UserID, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")
	TripID, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")

	trip := model.Trip{
		TripID: TripID,
		UserID: UserID,
	}

	if mockErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", mockErr)
	}

	rows := sqlmock.NewRows([]string{"ID"}).AddRow(UserID.Bytes())

	mock.ExpectQuery("INSERT INTO trips").WithArgs(UserID).WillReturnRows(rows)
	if _, err := model.CreateTrip(trip); err != nil {
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

	_, err := model.GetTripIDsByUserID(ID)

	if err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

//with TRIP_ID
func TestGetTrip(t *testing.T) {
	originalDB := database.DB
	database.DB, mock, mockErr = sqlmock.New()
	defer func() { database.DB = originalDB }()

	TripID, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")
	tests := []TripsTestCase{
		{
			name:            "GetTripsByTripId_200",
			mockedTripError: nil,
			expectedTripId:  TripID,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			model.GetEventsByTrip = func(tripID uuid.UUID) ([]model.Event, error) {
				return []model.Event{}, nil
			}
			model.GetFlightsByTrip = func(tripID uuid.UUID) ([]model.Flight, error) {
				return []model.Flight{}, nil
			}
			model.GetMuseumsByTrip = func(trip_id uuid.UUID) ([]model.Museum, error) {
				return []model.Museum{}, nil
			}
			model.GetRestaurantsFromTrip = func(tripsID uuid.UUID) ([]model.Restaurant, error) {
				return []model.Restaurant{}, nil
			}
			model.GetHotelsByTrip = func(tripID uuid.UUID) ([]model.Hotel, error) {
				return []model.Hotel{}, nil
			}
			model.GetTrainsFromTrip = func(tripsID uuid.UUID) ([]model.Train, error) {
				return []model.Train{}, nil
			}
			model.GetTripIDsByUserID = func(id uuid.UUID) ([]uuid.UUID, error) {
				return nil, nil
			}
			testTrip, _ := model.GetTrip(TripID)

			fmt.Println(testTrip)
		})
	}
}
