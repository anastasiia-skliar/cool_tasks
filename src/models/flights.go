package models

import (
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/satori/go.uuid"
	"net/url"
	"time"
)

const (
	addFlightToTrip = "INSERT INTO trips_flights (flight_id, trip_id) VALUES ($1, $2)"
	getFlightByTrip = "SELECT flights.* FROM flights INNER JOIN trips_flights ON flights.id=trips_flights.flight_id AND trips_flights.trip_id=$1"
)

//Flight is a representation of Flight table in DB
type Flight struct {
	ID            uuid.UUID
	DepartureCity string
	DepartureTime time.Time
	DepartureDate time.Time
	ArrivalCity   string
	ArrivalTime   time.Time
	ArrivalDate   time.Time
	Price         int
}

//AddFlightToTrip adds Flight to Trip
var AddFlightToTrip = func(flightID uuid.UUID, tripID uuid.UUID) error {
	_, err := database.DB.Exec(addFlightToTrip, flightID, tripID)
	return err
}

//GetFlightsByTrip gets Flights from Trip by tripID
var GetFlightsByTrip = func(tripID uuid.UUID) ([]Flight, error) {
	rows, err := database.DB.Query(getFlightByTrip, tripID)
	if err != nil {
		return nil, err
	}

	flights := make([]Flight, 0)

	for rows.Next() {
		var f Flight
		if err := rows.Scan(&f.ID, &f.DepartureCity, &f.DepartureTime, &f.DepartureDate, &f.ArrivalCity, &f.ArrivalDate, &f.ArrivalTime, &f.Price); err != nil {
			return nil, err
		}
		flights = append(flights, f)
	}
	return flights, nil
}

//GetFlights gets Flights from Trip by incoming request
var GetFlights = func(params url.Values) ([]Flight, error) {

	stringArgs := []string{"departure_city", "arrival_city"}
	numberArgs := []string{"price", "departure_time", "arrival_time", "departure_date", "arrival_date"}
	request, args, err := SQLGenerator("flights", stringArgs, numberArgs, params)
	if err != nil {
		return nil, err
	}
	rows, err := database.DB.Query(request, args...)
	if err != nil {
		return nil, err
	}

	flightsArray := make([]Flight, 0)

	for rows.Next() {
		var f Flight
		if err := rows.Scan(&f.ID, &f.DepartureCity, &f.DepartureTime, &f.DepartureDate, &f.ArrivalCity, &f.ArrivalDate, &f.ArrivalTime, &f.Price); err != nil {
			return nil, err
		}
		flightsArray = append(flightsArray, f)
	}
	return flightsArray, nil
}
