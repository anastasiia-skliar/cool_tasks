package museums

import (
	"github.com/satori/go.uuid"
	. "github.com/Nastya-Kruglikova/cool_tasks/src/database"
)

const (
	getMuseums      = "SELECT * FROM museums"
	getMuseumByCity = "SELECT * FROM museums WHERE location = $1"
	addMuseumToTrip = "INSERT INTO trips_museums (museum_id, trip_id) VALUES ($1, $2) RETURNING id"
	getMuseumsByTrip = "SELECT * FROM museums INNER JOIN trips_museums ON museums.id=trips_museums.museum_id AND trips_museums.trip_id=$1"
)

type Museum struct {
	ID          uuid.UUID
	name        string
	location    string
	price       int
	opened_at   int
	closed_at   int
	museum_type string
	info        string
}

var GetMuseums = func() ([]Museum, error) {
	rows, err := DB.Query(getMuseums)
	if err != nil {
		return []Museum{}, err
	}

	museums := make([]Museum, 0)

	for rows.Next() {
		var m Museum
		if err := rows.Scan(&m.ID, &m.name, &m.location, &m.price, &m.opened_at, &m.closed_at, &m.museum_type, &m.info); err != nil {
			return []Museum{}, err
		}
		museums = append(museums, m)
	}
	return museums, nil
}

var GetMuseumsByCity = func(city string) ([]Museum, error) {
	rows, err := DB.Query(getMuseumByCity)
	if err != nil {
		return []Museum{}, err
	}

	museums := make([]Museum, 0)

	for rows.Next() {
		var m Museum
		if err := rows.Scan(&m.ID, &m.name, &m.location, &m.price, &m.opened_at, &m.closed_at, &m.museum_type, &m.info); err != nil {
			return []Museum{}, err
		}
		museums = append(museums, m)
	}
	return museums, nil
}

var AddMuseumToTrip = func(museum_id uuid.UUID, trip_id uuid.UUID) (uuid.UUID, error){
	var id uuid.UUID
	err := DB.QueryRow(addMuseumToTrip, museum_id, trip_id).Scan(&id)

	return id, err
}

var GetMuseumsByTrip = func(trip_id uuid.UUID) ([]Museum, error) {
	rows, err := DB.Query(getMuseumsByTrip, trip_id)
	if err != nil {
		return []Museum{}, err
	}
	museums := make([]Museum, 0)

	for rows.Next() {
		var m Museum
		if err := rows.Scan(&m.ID, &m.name, &m.location, &m.price, &m.opened_at, &m.closed_at, &m.museum_type, &m.info); err != nil {
			return []Museum{}, err
		}
		museums = append(museums, m)
	}
	return museums, nil
}
