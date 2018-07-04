package museums

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/satori/go.uuid"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
)

var mock sqlmock.Sqlmock
var err error

func TestGetMuseums(t *testing.T) {
	originalDB := database.DB
	database.DB, mock, err = sqlmock.New()
	defer func() { database.DB = originalDB }()

	MuseumId, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")

	var expects = []Museum{
		{
			ID:          MuseumId,
			name:        "Ermitage",
			location:    "Peterburg",
			price:       1111,
			opened_at:   1,
			closed_at:   2,
			museum_type: "Gallery",
			info:        "Cool",
		},
		{
			ID:          MuseumId,
			name:        "Luvre",
			location:    "Paris",
			price:       1110,
			opened_at:   1,
			closed_at:   2,
			museum_type: "Gallery",
			info:        "Cool",
		},
	}

	rows := sqlmock.NewRows([]string{"ID", "name", "location", "price", "opened_at", "closed_at", "museum_type", "additional_info"}).

		AddRow(MuseumId.Bytes(), "Ermitage", "Peterburg", 1111, 1, 2, "Gallery", "Cool").AddRow(MuseumId.Bytes(), "Luvre", "Paris", 1110, 1, 2, "Gallery", "Cool")

	mock.ExpectQuery("SELECT (.+) FROM museums").WillReturnRows(rows)

	result, err := GetMuseums()

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

