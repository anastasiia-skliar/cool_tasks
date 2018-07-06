package models

import (
	"github.com/satori/go.uuid"
	. "github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"time"
)

const (
	getFlights         = "SELECT * FROM flights"
	getByCity          = "SELECT * FROM flights WHERE departure_city=$1 AND arrival_city=$2"
	getByDepartureTime = "SELECT * FROM flights WHERE departure_time BETWEEN $1 AND $2"
	getByArrivalTime   = "SELECT * FROM flights WHERE arrival_time BETWEEN $1 AND $2"
	getByPrice         = "SELECT * FROM flights WHERE price BETWEEN $1 AND $2"
	addToTrip          = "INSERT INTO trips_flights (flight_id, trip_id) VALUES ($1, $2)"
	getByDate          = "SELECT * FROM flights WHERE departure_date = $1 AND arrival_date = $2"
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

var GetFlights = func() ([]Flights, error) {
	rows, err := DB.Query(getFlights)
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

var GetByCity = func(departureCity, arrivalCity string) ([]Flights, error) {
	rows, err := DB.Query(getByCity, departureCity, arrivalCity)
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

var GetByDepartureTime = func(from, to time.Time) ([]Flights, error) {
	rows, err := DB.Query(getByDepartureTime, from, to)
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

var GetByArrivalTime = func(from, to time.Time) ([]Flights, error) {
	rows, err := DB.Query(getByArrivalTime, from, to)
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

var GetByPrice = func(from, to int) ([]Flights, error) {
	rows, err := DB.Query(getByPrice, from, to)
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

var AddToTrip = func(flightID uuid.UUID, tripID uuid.UUID) (error) {
	_, err := DB.Exec(addToTrip, flightID, tripID)
	return err
}

var GetByDate = func(departureDate, arrivalDate time.Time) ([]Flights, error) {
	rows, err := DB.Query(getByDate, departureDate, arrivalDate)
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