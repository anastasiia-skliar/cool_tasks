package models

import (
	"github.com/satori/go.uuid"
	. "github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"time"
)

const (
	addToTrip          = "INSERT INTO trips_flights (flight_id, trip_id) VALUES ($1, $2)"
	getByTrip          = "SELECT * FROM flights INNER JOIN trips_flights ON flights.id=trips_flights.flight_id AND trips_flights.trip_id=$1"
)

type Flights struct {
	ID            uuid.UUID
	departureCity string
	departureTime time.Time
	departureDate time.Time
	arrivalCity   string
	arrivalTime   time.Time
	arrivalDate   time.Time
	price         int
}

var AddToTrip = func(flightID uuid.UUID, tripID uuid.UUID) (error) {
	_, err := DB.Exec(addToTrip, flightID, tripID)
	return err
}

var GetByTrip = func(tripID uuid.UUID) ([]Flights, error) {
	rows, err := DB.Query(getByTrip, tripID)
	if err != nil {
		return []Flights{}, err
	}

	flights := make([]Flights, 0)

	for rows.Next() {
		var m Flights
		if err := rows.Scan(&m.ID, &m.departureCity, &m.departureTime, &m.departureDate, &m.arrivalCity, &m.arrivalDate, &m.arrivalTime, &m.price); err != nil {
			return []Flights{}, err
		}
		flights = append(flights, m)
	}
	return flights, nil
}

var GetByRequest = func(request string) ([]Flights, error) {
	rows, err := DB.Query(request)
	if err != nil {
		return []Flights{}, err
	}

	flights := make([]Flights, 0)

	for rows.Next() {
		var m Flights
		if err := rows.Scan(&m.ID, &m.departureCity, &m.departureTime, &m.departureDate, &m.arrivalCity, &m.arrivalDate, &m.arrivalTime, &m.price); err != nil {
			return []Flights{}, err
		}
		flights = append(flights, m)
	}
	return flights, nil
}