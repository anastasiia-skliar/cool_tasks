package models

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/satori/go.uuid"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"net/url"
	"testing"
	"time"
)

var eventMockErr error

type eventsTestCase struct {
	name string
	url  string
}

func TestGetEventsByRequest(t *testing.T) {
	testCase := []eventsTestCase{
		{"test with ID param",
			"localhost:8080/v1/events?id=00000000-0000-0000-0000-000000000001",
		},
		{"test with date price #1",
			"localhost:8080/v1/events?date=12:00:00&price=50",
		},
		{"test with 2 price params",
			"localhost:8080/v1/events?&price=50&price=300",
		},
		{
			"test with town, title, category",
			"localhost:8080/v1/events?&town=Kiev&title=Careerday&category=work",
		},
		{
			"test with 2 town, title, category",
			"localhost:8080/v1/events?&town=Kiev&town=Lviv&title=Careerday&category=work",
		},
	}

	originalDB := database.DB
	database.DB, mock, eventMockErr = sqlmock.New()
	defer func() { database.DB = originalDB }()
	testTime, _ := time.Parse("15:04:05", "12:00:00")
	EventID, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")

	var expects = []Event{
		{
			ID:       EventID,
			Title:    "Careerday",
			Category: "work",
			Town:     "Kiev",
			Date:     testTime,
			Price:    50,
		},
		{
			ID:       EventID,
			Title:    "ProjectX",
			Category: "entertaiment",
			Town:     "Lviv",
			Date:     testTime,
			Price:    300,
		},
	}

	for _, tc := range testCase {
		rawUrl, _ := url.Parse(tc.url)
		params := rawUrl.Query()

		if eventMockErr != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", eventMockErr)
		}

		rows := sqlmock.NewRows([]string{"ID", "Title", "Category", "Town", "Date", "Price"}).
			AddRow(EventID.Bytes(), "Careerday", "work", "Kiev", testTime, 50).
			AddRow(EventID.Bytes(), "ProjectX", "entertaiment", "Lviv", testTime, 300)

		mock.ExpectQuery("SELECT (.+) FROM events").WillReturnRows(rows)

		result, err := GetEvents(params)

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

func TestAddEventToTrip(t *testing.T) {
	originalDB := database.DB
	database.DB, mock, eventMockErr = sqlmock.New()
	defer func() { database.DB = originalDB }()

	TripID, _ := uuid.FromString("00000000-0000-0000-0000-000000000002")
	EventID, _ := uuid.FromString("00000000-0000-0000-0000-000000000003")

	if eventMockErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", eventMockErr)
	}

	mock.ExpectExec("INSERT INTO trips_events").WithArgs(EventID, TripID).WillReturnResult(sqlmock.NewResult(1, 1))
	if err := AddEventToTrip(EventID, TripID); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetEventsByTrip(t *testing.T) {
	originalDB := database.DB
	database.DB, mock, eventMockErr = sqlmock.New()
	defer func() { database.DB = originalDB }()

	testTime, _ := time.Parse("15:04:05", "12:00:00")
	EventID, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")

	expected := Event{

		EventID,
		"Careerday",
		"work",
		"Kiev",
		testTime,
		50,
	}

	if eventMockErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", eventMockErr)
	}

	rows := sqlmock.NewRows([]string{"EventID", "Title", "Category", "Town", "Date", "Price"}).
		AddRow(EventID.Bytes(), "Careerday", "work", "Kiev", testTime, 50)

	mock.ExpectQuery("SELECT (.+) FROM events INNER JOIN trips_events ON events.id=trips_events.event_id AND trips_events.trip_id=\\$1").WithArgs(expected.ID).WillReturnRows(rows)

	result, err := GetEventsByTrip(EventID)

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
