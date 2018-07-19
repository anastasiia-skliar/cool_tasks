package models

import (
	. "github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/satori/go.uuid"
	"net/url"
	"time"
)

const (
	addMuseumToTrip  = "INSERT INTO trips_museums (museum_id, trip_id) VALUES ($1, $2)"
	getMuseumsByTrip = "SELECT museums.* FROM museums INNER JOIN trips_museums ON museums.id=trips_museums.museum_id AND trips_museums.trip_id=$1"
)

type Museum struct {
	ID         uuid.UUID
	Name       string
	Location   string
	Price      int
	OpenedAt   time.Time
	ClosedAt   time.Time
	MuseumType string
	Info       string
}

var AddMuseumToTrip = func(museum_id uuid.UUID, trip_id uuid.UUID) error {
	_, err := DB.Exec(addMuseumToTrip, museum_id, trip_id)
	return err
}

var GetMuseumsByTrip = func(trip_id uuid.UUID) ([]Museum, error) {
	rows, err := DB.Query(getMuseumsByTrip, trip_id)
	if err != nil {
		return nil, err
	}
	museums := make([]Museum, 0)

	for rows.Next() {
		var m Museum
		if err := rows.Scan(&m.ID, &m.Name, &m.Location, &m.Price, &m.OpenedAt, &m.ClosedAt, &m.MuseumType, &m.Info); err != nil {
			return nil, err
		}
		museums = append(museums, m)
	}
	return museums, nil
}

var GetMuseumsByRequest = func(params url.Values) ([]Museum, error) {
	stringArgs := []string{"name", "location", "museum_type"}
	numberArgs := []string{"price", "opened_at", "closed_at"}
	request, args, err := SQLGenerator("museums", stringArgs, numberArgs, params)
	if err != nil {
		return nil, err
	}
	rows, err := DB.Query(request, args...)
	if err != nil {
		return nil, err
	}

	museumsArray := make([]Museum, 0)

	for rows.Next() {
		var m Museum
		if err := rows.Scan(&m.ID, &m.Name, &m.Location, &m.Price, &m.OpenedAt, &m.ClosedAt, &m.MuseumType, &m.Info); err != nil {
			return nil, err
		}
		museumsArray = append(museumsArray, m)
	}
	return museumsArray, nil
}
