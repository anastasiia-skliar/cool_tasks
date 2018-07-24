package models

import (
	"fmt"
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
	if err := AddRestaurantToTrip(restaurantID, tripID); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetRestaurantByTripID(t *testing.T) {
	originalDB := database.DB
	database.DB, mock, flightMockErr = sqlmock.New()
	defer func() { database.DB = originalDB }()

	ID, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")

	expected := []Restaurant{
		{
			ID,
			"Крива липа",
			"Lviv",
			4,
			3,
			"Кулінарна студія «Крива Липа» – це авторська кухня без ГМО. Справжні кулінарні шедеври тільки найкращої якості та зі свіжих продуктів від знаних майстрів своєї справи",
		},
		{
			ID,
			"Крива липа",
			"Lviv",
			4,
			3,
			"Кулінарна студія «Крива Липа» – це авторська кухня без ГМО. Справжні кулінарні шедеври тільки найкращої якості та зі свіжих продуктів від знаних майстрів своєї справи",
		},
	}
	if flightMockErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", flightMockErr)
	}

	rows := sqlmock.NewRows([]string{"ID", "name", "location", "stars", "prices", "description"}).
		AddRow(ID.Bytes(), "Крива липа", "Lviv", 4, 3, "Кулінарна студія «Крива Липа» – це авторська кухня без ГМО. Справжні кулінарні шедеври тільки найкращої якості та зі свіжих продуктів від знаних майстрів своєї справи")

	mock.ExpectQuery("SELECT (.+) FROM restaurants").WithArgs(ID).WillReturnRows(rows)

	result, err := GetRestaurantsFromTrip(ID)
	fmt.Println(result)

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

func TestRequestGenerator(t *testing.T) {
	generatorTests := []reqGenTestCases{
		{
			"SELECT * FROM restaurants WHERE ID = ",
			"ID",
			"SELECT * FROM restaurants WHERE ID = $1",
		},
		{
			"SELECT * FROM restaurants",
			"",
			"SELECT * FROM restaurants",
		},
	}
	for _, tc := range generatorTests {
		var request = ""
		if tc.paramName == "" {
			request = recGen()
		} else {
			request = recGen(tc.paramName)
		}
		if request != tc.expected {
			t.Error("Expected:", tc.expected, "Was:", request)
		}
	}
}
