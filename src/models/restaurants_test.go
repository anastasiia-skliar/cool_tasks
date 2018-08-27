package models

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/satori/go.uuid"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
)

var restaurantsMockErr error

type restaurantsTestCase struct {
	name string
	url  string
}
type reqGenTestCases struct {
	name      string
	paramName string
	expected  string
}

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

//func TestGetRestaurantByTripID(t *testing.T) {
//	originalDB := database.DB
//	database.DB, mock, flightMockErr = sqlmock.New()
//	defer func() { database.DB = originalDB }()
//
//	ID, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")
//
//	expected := []Restaurant{
//		{
//			ID,
//			"Kryva Lypa",
//			"Lviv",
//			4,
//			3,
//			"Some info 1",
//		},
//		{
//			ID,
//			"Kryva Lypa",
//			"Lviv",
//			4,
//			3,
//			"Some info 1",
//		},
//	}
//	if flightMockErr != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", flightMockErr)
//	}
//
//	rows := sqlmock.NewRows([]string{"ID", "name", "location", "stars", "prices", "description"}).
//		AddRow(ID.Bytes(), expected[0].Name, expected[0].Location, expected[0].Stars, expected[0].Prices, expected[0].Description)
//
//	mock.ExpectQuery("SELECT (.+) FROM restaurants").WithArgs(ID).WillReturnRows(rows)
//
//	result, err := GetRestaurantsFromTrip(ID)
//	fmt.Println(result)
//
//	if err != nil {
//		t.Errorf("error was not expected while updating stats: %s", err)
//	}
//	if err := mock.ExpectationsWereMet(); err != nil {
//		t.Errorf("there were unfulfilled expectations: %s", err)
//	}
//
//	for i := 0; i < len(result); i++ {
//		if expected[i] != result[i] {
//			t.Error("Expected:", expected[i], "Was:", result[i])
//		}
//	}
//}
