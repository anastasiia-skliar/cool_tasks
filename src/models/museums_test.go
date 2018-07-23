package models

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/satori/go.uuid"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"net/url"
	"testing"
	"time"
)

var museumMockErr error

type museumsTestCase struct {
	name string
	url  string
}

func TestGetMuseumsByRequest(t *testing.T) {
	testCases := []museumsTestCase{
		{
			"test with ID ",
			"/v1/museums?id=00000000-0000-0000-0000-000000000001",
		},
		{
			"test with price ,opened_at,closed_at",
			"/v1/museums?price=1111&opened_at=12:00:00&closed_at=12:00:00",
		},
		{
			"name",
			"/v1/museums?name=Ermitage",
		},
		{
			"test with 2 prices",
			"/v1/museums?price=1111&price=1110",
		},
		{
			"  test with 2 museum_type",
			"/v1/museums?location=Peterburg&location=Paris",
		},
	}
	originalDB := database.DB
	database.DB, mock, museumMockErr = sqlmock.New()
	defer func() { database.DB = originalDB }()

	for _, tc := range testCases {
		testTime, _ := time.Parse("15:04:05", "12:00:00")
		MuseumID, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")

		var expects = []Museum{
			{
				ID:         MuseumID,
				Name:       "Ermitage",
				Location:   "Peterburg",
				Price:      1111,
				OpenedAt:   testTime,
				ClosedAt:   testTime,
				MuseumType: "Gallery",
				Info:       "Cool",
			},
			{
				ID:         MuseumID,
				Name:       "Luvre",
				Location:   "Paris",
				Price:      1110,
				OpenedAt:   testTime,
				ClosedAt:   testTime,
				MuseumType: "Gallery",
				Info:       "Cool",
			},
		}

		if museumMockErr != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", museumMockErr)
		}

		rows := sqlmock.NewRows([]string{"ID", "Name", "Location", "Price", "OpenedAt", "ClosedAt", "MuseumType", "additional_info"}).
			AddRow(MuseumID.Bytes(), "Ermitage", "Peterburg", 1111, testTime, testTime, "Gallery", "Cool").
			AddRow(MuseumID.Bytes(), "Luvre", "Paris", 1110, testTime, testTime, "Gallery", "Cool")

		mock.ExpectQuery("SELECT (.+) FROM museums").WillReturnRows(rows)
		rawUrl, _ := url.Parse(tc.url)
		params := rawUrl.Query()

		result, err := GetMuseumsByRequest(params)

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

func TestAddMuseumToTrip(t *testing.T) {
	originalDB := database.DB
	database.DB, mock, museumMockErr = sqlmock.New()
	defer func() { database.DB = originalDB }()

	TripID, _ := uuid.FromString("00000000-0000-0000-0000-000000000002")
	MuseumID, _ := uuid.FromString("00000000-0000-0000-0000-000000000003")

	if museumMockErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", museumMockErr)
	}

	mock.ExpectExec("INSERT INTO trips_museums").WithArgs(MuseumID, TripID).WillReturnResult(sqlmock.NewResult(1, 1))
	if err := AddMuseumToTrip(MuseumID, TripID); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetMuseumByTrip(t *testing.T) {
	originalDB := database.DB
	database.DB, mock, museumMockErr = sqlmock.New()
	defer func() { database.DB = originalDB }()

	testTime, _ := time.Parse("15:04:05", "12:00:00")
	MuseumID, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")

	expected := Museum{

		MuseumID,
		"Louvre",
		"Paris",
		1111,
		testTime,
		testTime,
		"History",
		"Cool",
	}

	if museumMockErr != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", museumMockErr)
	}

	rows := sqlmock.NewRows([]string{"MuseumID", "Name", "Location", "Price", "OpenedAt", "ClosedAt", "MuseumType", "additional_info"}).
		AddRow(MuseumID.Bytes(), "Louvre", "Paris", 1111, testTime, testTime, "History", "Cool")

	mock.ExpectQuery("SELECT (.+) FROM museums INNER JOIN trips_museums ON museums.id=trips_museums.museum_id AND trips_museums.trip_id=\\$1").WithArgs(expected.ID).WillReturnRows(rows)

	result, err := GetMuseumsByTrip(MuseumID)

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
