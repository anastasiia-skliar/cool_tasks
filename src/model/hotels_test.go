package model

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/satori/go.uuid"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"net/url"
	"testing"
)

var (
	mockErr error
)

type hotelsTestCase struct {
	name string
	url  string
}

func TestAddHotelToTrip(t *testing.T) {
	originalDB := database.DB
	database.DB, mock, mockErr = sqlmock.New()
	defer func() { database.DB = originalDB }()

	hotelID, _ := uuid.FromString("00000000-0000-0000-0000-000000000002")
	tripID, _ := uuid.FromString("00000000-0000-0000-0000-000000000003")

	if mockErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", mockErr)
	}

	mock.ExpectExec("INSERT INTO trips_hotel").WithArgs(hotelID, tripID).WillReturnResult(sqlmock.NewResult(1, 1))
	if err := AddHotelToTrip(hotelID, tripID); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}

}

func TestGetHotelsByTrip(t *testing.T) {
	originalDB := database.DB
	database.DB, mock, mockErr = sqlmock.New()
	defer func() { database.DB = originalDB }()

	if mockErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", mockErr)
	}

	ID, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")

	expects := []Hotel{
		{
			ID:        ID,
			Name:      "Hotel Ukraine",
			Class:     3,
			Capacity:  1000,
			RoomsLeft: 218,
			Floors:    18,
			MaxPrice:  3200,
			CityName:  "Kyiv",
			Address:   "Vulytsya Instytutsʹka 4",
		},
		{
			ID:        ID,
			Name:      "Lviv",
			Class:     4,
			Capacity:  1450,
			RoomsLeft: 200,
			Floors:    9,
			MaxPrice:  3480,
			CityName:  "Lviv",
			Address:   "Prospect V. Chornovil, 7",
		},
	}

	rows := sqlmock.NewRows([]string{"ID", "name", "class", "capacity", "rooms_left",
		"floors", "max_price", "city_name", "address"}).
		AddRow(ID.Bytes(), "Hotel Ukraine", 3, 1000, 218, 18,
			3200, "Kyiv", "Vulytsya Instytutsʹka 4")

	mock.ExpectQuery("SELECT (.+) FROM hotels INNER JOIN trips_hotels ON hotels.id=trips_hotels.hotels_id AND trips_hotels.trip_id").WithArgs(ID).WillReturnRows(rows)

	result, err := GetHotelsByTrip(ID)

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

func TestGetHotelByRequest(t *testing.T) {
	testCases := []hotelsTestCase{
		{"test with zero parameters",
			"localhost:8080/v1/hotels",
		},
		{
			"test with all params",
			"localhost:8080/v1/hotels?id=00000000-0000-0000-0000-000000000001&name=Lviv&class=4&capacity=1450&rooms_left=200" +
				"&max_price=3480uah&city_name=Lviv&address=Prospect V. Chornovil, 7",
		},
		{
			"test with 2 max_prices",
			"localhost:8080/v1/hotels?max_price=3200uah&max_price=3480uah",
		},
	}

	originalDB := database.DB
	database.DB, mock, mockErr = sqlmock.New()
	defer func() { database.DB = originalDB }()

	for _, tc := range testCases {
		rawUrl, _ := url.Parse(tc.url)
		params := rawUrl.Query()
		ID, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")

		expects := []Hotel{
			{
				ID:        ID,
				Name:      "Hotel Ukraine",
				Class:     3,
				Capacity:  1000,
				RoomsLeft: 218,
				Floors:    18,
				MaxPrice:  3200,
				CityName:  "Kyiv",
				Address:   "Vulytsya Instytutsʹka 4",
			},
			{
				ID:        ID,
				Name:      "Lviv",
				Class:     4,
				Capacity:  1450,
				RoomsLeft: 200,
				Floors:    9,
				MaxPrice:  3480,
				CityName:  "Lviv",
				Address:   "Prospect V. Chornovil, 7",
			},
		}
		if mockErr != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", mockErr)
		}
		rows := sqlmock.NewRows([]string{"ID", "name", "class", "capacity", "rooms_left",
			"floors", "max_price", "city_name", "address"}).
			AddRow(ID.Bytes(), "Hotel Ukraine", 3, 1000, 218, 18,
				3200, "Kyiv", "Vulytsya Instytutsʹka 4").AddRow(ID.Bytes(), "Lviv", 4, 1450, 200, 9,
			3480, "Lviv", "Prospect V. Chornovil, 7")

		mock.ExpectQuery("SELECT (.+) FROM hotels").WillReturnRows(rows)

		result, err := GetHotels(params)

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

}
