package museums

import (
	"github.com/satori/go.uuid"
	. "github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"fmt"
)

const (
	getMuseums       = "SELECT * FROM museums"
	getMuseumByCity  = "SELECT * FROM museums WHERE Location = ?"
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

var GetMuseums = func() ([]Museum, error) {
	rows, err := DB.Query(getMuseums)
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

var GetMuseumsByCity = func(city string) ([]Museum, error) {
	rows, err := DB.Query(getMuseumByCity, city)
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
	fmt.Println(museums)
	return museums, nil
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

var GetMuseumsByRequest = func(request string) ([]Museum, error) {
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
