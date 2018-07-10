package models

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/satori/go.uuid"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
	"net/http"
)


func TestGetMuseumsByRequest(t *testing.T) {
	originalDB := database.DB
	database.DB, mock, err = sqlmock.New()
	defer func() { database.DB = originalDB }()

	MuseumId, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")

	var expects = []Museum{
		{
			ID:         MuseumId,
			Name:       "Ermitage",
			Location:   "Peterburg",
			Price:      1111,
			OpenedAt:   1,
			ClosedAt:   2,
			MuseumType: "Gallery",
			Info:       "Cool",
		},
		{
			ID:         MuseumId,
			Name:       "Luvre",
			Location:   "Paris",
			Price:      1110,
			OpenedAt:   1,
			ClosedAt:   2,
			MuseumType: "Gallery",
			Info:       "Cool",
		},
	}

	rows := sqlmock.NewRows([]string{"ID", "Name", "Location", "Price", "OpenedAt", "ClosedAt", "MuseumType", "additional_info"}).

		AddRow(MuseumId.Bytes(), "Ermitage", "Peterburg", 1111, 1, 2, "Gallery", "Cool").AddRow(MuseumId.Bytes(), "Luvre", "Paris", 1110, 1, 2, "Gallery", "Cool")

	mock.ExpectQuery("SELECT (.+) FROM museums").WillReturnRows(rows)

	req, _ := http.NewRequest(http.MethodGet, "/v1/museums", nil)

	result, err := GetMuseumsByRequest(req.URL.Query())

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

func TestAddMuseumToTrip(t *testing.T) {
	originalDB := database.DB
	database.DB, mock, err = sqlmock.New()
	defer func() { database.DB = originalDB }()

	TripId, _ := uuid.FromString("00000000-0000-0000-0000-000000000002")
	MuseumId, _ := uuid.FromString("00000000-0000-0000-0000-000000000003")

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectExec("INSERT INTO trips_museums").WithArgs(MuseumId, TripId).WillReturnResult(sqlmock.NewResult(1, 1))
	if err := AddMuseumToTrip(MuseumId, TripId); err != nil {
		t.Errorf("error was not expected while updating stats: %s", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetMuseumByTrip(t *testing.T) {
	originalDB := database.DB
	database.DB, mock, err = sqlmock.New()
	defer func() { database.DB = originalDB }()

	MuseumId, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")

	expected := Museum{

		MuseumId,
		"Louvre",
		"Paris",
		1111,
		1,
		2,
		"History",
		"Cool",
	}

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	rows := sqlmock.NewRows([]string{"MuseumId", "Name", "Location", "Price", "OpenedAt", "ClosedAt", "MuseumType", "additional_info"}).
		AddRow(MuseumId.Bytes(), "Louvre", "Paris", 1111, 1, 2, "History", "Cool")

	mock.ExpectQuery("SELECT (.+) FROM museums INNER JOIN trips_museums ON museums.id=trips_museums.museum_id AND trips_museums.trip_id=\\$1").WithArgs(expected.ID).WillReturnRows(rows)

	result, err := GetMuseumsByTrip(MuseumId)

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
