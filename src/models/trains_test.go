package models_test

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/Nastya-Kruglikova/cool_tasks/src/models"
	"github.com/satori/go.uuid"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
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

//func TestGetTrains(t *testing.T) {
//	testCases := []trainsTestCase{
//		{"id departure_time departure_date arrival_time arrival_date price",
//			"localhost:8080/hello?id=00000000-0000-0000-0000-000000000001&departure_time=12:00:00" +
//				"&arrival_time=15:00:00&arrival_date=2018-07-21&price=200uah",
//		},
//		{
//			"Departure city & arrival City",
//			"localhost:8080/hello?departure_city=Lviv&arrival_city=Kyiv",
//		},
//		{
//			"price > 200 & price <250",
//			"localhost:8080/hello?price=200uah&price=250uah",
//		},
//		{"test with zero parameters",
//			"localhost:8080/hello",
//		},
//	}
//	originalDB := database.DB
//	database.DB, mock, mockErr = sqlmock.New()
//	defer func() { database.DB = originalDB }()
//
//	for _, tc := range testCases {
//		rawUrl, _ := url.Parse(tc.url)
//		params := rawUrl.Query()
//		ID, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")
//
//		departureTime, _ := time.Parse("15:04:05", "12:00:00")
//		departureDate, _ := time.Parse("2006-01-02", "2018-07-21")
//		arrivalTime, _ := time.Parse("15:04:05", "15:00:00")
//		arrivalDate, _ := time.Parse("2006-01-02", "2018-07-21")
//
//		expects := []models.Train{
//			{
//				ID:            ID,
//				DepartureTime: departureTime,
//				DepartureDate: departureDate,
//				ArrivalTime:   arrivalTime,
//				ArrivalDate:   arrivalDate,
//				DepartureCity: "Lviv",
//				ArrivalCity:   "Kyiv",
//				TrainType:     "el",
//				CarType:       "coupe",
//				Price:         "200uah",
//			},
//			{
//				ID:            ID,
//				DepartureTime: departureTime,
//				DepartureDate: departureDate,
//				ArrivalTime:   arrivalTime,
//				ArrivalDate:   arrivalDate,
//				DepartureCity: "Lviv",
//				ArrivalCity:   "Kyiv",
//				TrainType:     "el",
//				CarType:       "coupe",
//				Price:         "250uah",
//			},
//		}
//
//		if mockErr != nil {
//			t.Fatalf("an error '%s' was not expected when opening a stub database connection", mockErr)
//		}
//
//		rows := sqlmock.NewRows([]string{"ID", "departureTime", "departureDate", "arrivalTime", "arrivalDate",
//			"departure_city", "arrival_city", "train_type", "car_type", "price"}).
//			AddRow(ID.Bytes(), departureTime, departureDate, arrivalTime, arrivalDate,
//				"Lviv", "Kyiv", "el", "coupe", "200uah").AddRow(ID.Bytes(), departureTime, departureDate, arrivalTime, arrivalDate,
//			"Lviv", "Kyiv", "el", "coupe", "250uah")
//
//		mock.ExpectQuery("SELECT (.+) FROM trains").WillReturnRows(rows)
//		models.SQLGenerator = originalGenerator
//		result, err := models.GetData(params)
//
//		if err != nil {
//			t.Errorf("error was not expected while updating stats: %s", err)
//		}
//		if err := mock.ExpectationsWereMet(); err != nil {
//			t.Errorf("there were unfulfilled expectations: %s", err)
//		}
//
//		for i := 0; i < len(result); i++ {
//			if expects[i] != result[i] {
//				t.Error("Expected:", expects[i], "Was:", result[i])
//			}
//		}
//	}
//}

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
	if err := models.AddToTrip(trainID, tripID,models.Train{}); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

//func TestGetTrainFromTrip(t *testing.T) {
//	originalDB := database.DB
//	database.DB, mock, mockErr = sqlmock.New()
//	defer func() { database.DB = originalDB }()
//
//	if mockErr != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", mockErr)
//	}
//
//	departureTime, _ := time.Parse("15:04:05", "12:00:00")
//	departureDate, _ := time.Parse("2006-01-02", "2018-07-21")
//	arrivalTime, _ := time.Parse("15:04:05", "15:00:00")
//	arrivalDate, _ := time.Parse("2006-01-02", "2018-07-21")
//
//	ID, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")
//
//	expects := []models.Train{
//		{
//			ID:            ID,
//			DepartureTime: departureTime,
//			DepartureDate: departureDate,
//			ArrivalTime:   arrivalTime,
//			ArrivalDate:   arrivalDate,
//			DepartureCity: "Lviv",
//			ArrivalCity:   "Kyiv",
//			TrainType:     "el",
//			CarType:       "coupe",
//			Price:         "200uah",
//		},
//		{
//			ID:            ID,
//			DepartureTime: departureTime,
//			DepartureDate: departureDate,
//			ArrivalTime:   arrivalTime,
//			ArrivalDate:   arrivalDate,
//			DepartureCity: "Lviv",
//			ArrivalCity:   "Kharkiv",
//			TrainType:     "el",
//			CarType:       "coupe",
//			Price:         "250uah",
//		},
//	}
//
//	rows := sqlmock.NewRows([]string{"ID", "departure_time", "departure_date", "arrival_time", "arrival_date",
//		"departure_city", "arrival_city", "train_type", "car_type", "price"}).
//		AddRow(ID.Bytes(), departureTime, departureDate, arrivalTime, arrivalDate,
//			"Lviv", "Kyiv", "el", "coupe", "200uah")
//
//	mock.ExpectQuery("SELECT (.+) FROM trains INNER JOIN trips_trains ON trips_trains.train_id = trains.id AND trips_trains.trip_id").WithArgs(ID).WillReturnRows(rows)
//
//	result, err := models.GetTrainsFromTrip(ID)
//
//	if err != nil {
//		t.Errorf("error was not expected while updating stats: %s", err)
//	}
//	if err := mock.ExpectationsWereMet(); err != nil {
//		t.Errorf("there were unfulfilled expectations: %s", err)
//	}
//
//	for i := 0; i < len(result); i++ {
//		if expects[i] != result[i] {
//			t.Error("Expected:", expects[i], "Was:", result[i])
//		}
//	}
//}
