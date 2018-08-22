package models_test

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/Nastya-Kruglikova/cool_tasks/src/models"
	"github.com/satori/go.uuid"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"net/url"
	"testing"
	"time"
)

var (
	mock    sqlmock.Sqlmock
	mockErr error
)

type trainsTestCase struct {
	name string
	url  string
}

var originalGenerator = models.SQLGenerator

func TestGetTrains(t *testing.T) {
	testCases := []trainsTestCase{
		{"id departure arrival",
			"localhost:8080/hello?id=00000000-0000-0000-0000-000000000001&departure=2006-01-02 12:00:00" +
				"&arrival=2006-01-02 15:00:00&price=200",
		},
		{
			"Departure city & arrival City",
			"localhost:8080/hello?departure_city=Lviv&arrival_city=Kyiv",
		},
		{
			"price > 200 & price < 250",
			"localhost:8080/hello?price=200uah&price=250uah",
		},
		{"test with zero parameters",
			"localhost:8080/hello",
		},
	}
	originalDB := database.DB
	database.DB, mock, mockErr = sqlmock.New()
	defer func() { database.DB = originalDB }()

	for _, tc := range testCases {
		rawUrl, _ := url.Parse(tc.url)
		params := rawUrl.Query()
		ID, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")

		departure, _ := time.Parse("15:04:05", "12:00:00")
		arrival, _ := time.Parse("2006-01-02", "2018-07-21")

		expects := []models.Train{
			{
				ID:            ID,
				Departure:     departure,
				Arrival:       arrival,
				DepartureCity: "Lviv",
				ArrivalCity:   "Kyiv",
				TrainType:     "el",
				CarType:       "coupe",
				Price:         200,
			},
			{
				ID:            ID,
				Departure:     departure,
				Arrival:       arrival,
				DepartureCity: "Lviv",
				ArrivalCity:   "Kyiv",
				TrainType:     "el",
				CarType:       "coupe",
				Price:         250,
			},
		}

		if mockErr != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", mockErr)
		}

		rows := sqlmock.NewRows([]string{"ID", "departure", "arrival",
			"departure_city", "arrival_city", "train_type", "car_type", "price"}).
			AddRow(ID.Bytes(), departure, arrival,
				"Lviv", "Kyiv", "el", "coupe", 200).AddRow(ID.Bytes(), departure, arrival,
			"Lviv", "Kyiv", "el", "coupe", 250)

		mock.ExpectQuery("SELECT (.+) FROM trains").WillReturnRows(rows)
		models.SQLGenerator = originalGenerator
		result, err := models.GetTrains(params)

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

func TestSaveTrainToTrip(t *testing.T) {

	originalDB := database.DB
	database.DB, mock, mockErr = sqlmock.New()
	defer func() { database.DB = originalDB }()

	trainID, _ := uuid.FromString("00000000-0000-0000-0000-000000000002")
	tripID, _ := uuid.FromString("00000000-0000-0000-0000-000000000003")

	if mockErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", mockErr)
	}

	mock.ExpectExec("INSERT INTO trips_trains").WithArgs(trainID, tripID).WillReturnResult(sqlmock.NewResult(1, 1))
	if err := models.AddTrainToTrip(trainID, tripID); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetTrainFromTrip(t *testing.T) {
	originalDB := database.DB
	database.DB, mock, mockErr = sqlmock.New()
	defer func() { database.DB = originalDB }()

	if mockErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", mockErr)
	}

	departure, _ := time.Parse("2006-01-02 15:04:05", " 2006-01-02 12:00:00")
	arrival, _ := time.Parse("2006-01-02 15:04:05", "2006-01-02 15:00:00")

	ID, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")

	expects := []models.Train{
		{
			ID:            ID,
			Departure:     departure,
			Arrival:       arrival,
			DepartureCity: "Lviv",
			ArrivalCity:   "Kyiv",
			TrainType:     "el",
			CarType:       "coupe",
			Price:         200,
		},
		{
			ID:            ID,
			Departure:     departure,
			Arrival:       arrival,
			DepartureCity: "Lviv",
			ArrivalCity:   "Kharkiv",
			TrainType:     "el",
			CarType:       "coupe",
			Price:         250,
		},
	}

	rows := sqlmock.NewRows([]string{"ID", "departure", "arrival",
		"departure_city", "arrival_city", "train_type", "car_type", "price"}).
		AddRow(ID.Bytes(), departure, arrival, "Lviv", "Kyiv", "el", "coupe", 200)

	mock.ExpectQuery("SELECT (.+) FROM trains INNER JOIN trips_trains ON trips_trains.train_id = trains.id AND trips_trains.trip_id").WithArgs(ID).WillReturnRows(rows)

	result, err := models.GetTrainsFromTrip(ID)

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
