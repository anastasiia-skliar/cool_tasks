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

	trip := models.Trip{
		TripID: TripID,
		UserID: UserID,
	}

	if mockErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", mockErr)
	}

	rows := sqlmock.NewRows([]string{"ID"}).AddRow(UserID.Bytes())

	mock.ExpectQuery("INSERT INTO trips").WithArgs(UserID).WillReturnRows(rows)
	if _, err := models.AddTrip(trip); err != nil {
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

	_, err := models.GetTripIDsByUserID(ID)

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
	//testTime, _ := time.Parse("15:04:05", "12:00:00")
	tests := []TripsTestCase{
		{
			name:            "GetTripsByTripId_200",
			mockedTripError: nil,
			expectedTripId:  TripID,
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {

			models.GetEventsByTrip = func(tripID uuid.UUID) ([]models.Event, error) {
				return []models.Event{}, nil
			}
			models.GetFlightsByTrip = func(tripID uuid.UUID) ([]models.Flight, error) {
				return []models.Flight{}, nil
			}
			models.GetMuseumsByTrip = func(trip_id uuid.UUID) ([]models.Museum, error) {
				return []models.Museum{}, nil
			}
			models.GetRestaurantsFromTrip = func(tripsID uuid.UUID) ([]models.Restaurant, error) {
				return []models.Restaurant{}, nil
			}
			models.GetHotelsByTrip = func(tripID uuid.UUID) ([]models.Hotel, error) {
				return []models.Hotel{}, nil
			}
			models.GetTrainsFromTrip = func(tripsID uuid.UUID) ([]models.Train, error) {
				return []models.Train{}, nil
			}
			models.GetTripIDsByUserID = func(id uuid.UUID) ([]uuid.UUID, error) {
				return nil, nil
			}
			testTrip, _ := models.GetTrip(TripID)
			
			if testTrip.TripID != tc.expectedTripId {
				t.Errorf("Expected: %s", tc.name)
			}
		})
	}
}
