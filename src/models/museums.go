package models

import (
	"github.com/satori/go.uuid"
	. "github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"net/url"
	"strings"
	"errors"
)

const (
	addMuseumToTrip  = "INSERT INTO trips_museums (museum_id, trip_id) VALUES ($1, $2)"
	getMuseumsByTrip = "SELECT * FROM museums INNER JOIN trips_museums ON museums.id=trips_museums.museum_id AND trips_museums.trip_id=$1"
)

type Museum struct {
	ID         uuid.UUID
	Name       string
	Location   string
	Price      int
	OpenedAt   int
	ClosedAt   int
	MuseumType string
	Info       string
}

var AddMuseumToTrip = func(museum_id uuid.UUID, trip_id uuid.UUID) (error) {
	_, err := DB.Exec(addMuseumToTrip, museum_id, trip_id)
	return err
}

var GetMuseumsByTrip = func(trip_id uuid.UUID) ([]Museum, error) {
	rows, err := DB.Query(getMuseumsByTrip, trip_id)
	if err != nil {
		return []Museum{}, err
	}
	museums := make([]Museum, 0)

	for rows.Next() {
		var m Museum
		if err := rows.Scan(&m.ID, &m.Name, &m.Location, &m.Price, &m.OpenedAt, &m.ClosedAt, &m.MuseumType, &m.Info); err != nil {
			return []Museum{}, err
		}
		museums = append(museums, m)
	}
	return museums, nil
}

var GetMuseumsByRequest = func(params url.Values) ([]Museum, error) {
	request := "SELECT * FROM museums WHERE "
	count := 0
	validKeys := []string{"id", "name", "location", "price", "museum_type", "opened_at", "closed_at"}
	for key, value := range params {
		for _, keys := range validKeys {
			if key == keys {
				count++
			}
		}
		if count == 0 {

			return []Museum{}, errors.New("ERROR: Bad request")
		}
		switch key {
		case "name", "location", "museum_type":
			if len(value) > 1 {
				request += key + " IN ("
				for i, v := range value {
					v = "'" + v + "'"
					request += v
					if i < len(value)-1 {
						request += ", "
					}
				}
				request += ") AND "
			} else {
				value[0] = "'" + value[0] + "'"
				request += key + "=" + value[0] + " AND "
			}
		case "price", "opened_at", "closed_at":
			if len(value) == 2 {
				request += key + " BETWEEN " + value[0] + " AND " + value[1] + " AND "
			} else {
				request += key + "=" + value[0] + " AND "
			}
		default:
			request += key + "=" + value[0] + " AND "
		}

		count = 0
	}

	words := strings.Fields(request)

	if words[len(words)-1] == "AND" || words[len(words)-1] == "WHERE" {
		words[len(words)-1] = ""
	}

	request = strings.Join(words, " ")
	request+=";"
	rows, err := DB.Query(request)
	if err != nil {
		return []Museum{}, err
	}

	museums := make([]Museum, 0)

	for rows.Next() {
		var m Museum
		if err := rows.Scan(&m.ID, &m.Name, &m.Location, &m.Price, &m.OpenedAt, &m.ClosedAt, &m.MuseumType, &m.Info); err != nil {
			return []Museum{}, err
		}
		museums = append(museums, m)
	}
	return museums, nil
}
