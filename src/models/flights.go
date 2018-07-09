package models

import (
	"github.com/satori/go.uuid"
	. "github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"time"
	"net/url"
	"errors"
	sq "github.com/Masterminds/squirrel"
)

const (
	addToTrip = "INSERT INTO trips_flights (flight_id, trip_id) VALUES ($1, $2)"
	getByTrip = "SELECT * FROM flights INNER JOIN trips_flights ON flights.id=trips_flights.flight_id AND trips_flights.trip_id=$1"
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

var GetByRequest = func(params url.Values) ([]Flights, error) {

	var and sq.And
	var or sq.Or
	count := 0
	flights := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).Select("*").From("flights")
	validKeys := []string{"id", "departure_city", "departure_time", "departure_date", "arrival_city", "arrival_time", "arrival_date", "price"}
	for key, value := range params {
		for _, keys := range validKeys {
			if key == keys {
				count++
			}
		}
		if count == 0 {
			return []Flights{}, errors.New("ERROR: Bad request")
		}

		switch key {
		case "departure_city", "arrival_city":
			if len(value) > 1 {
				for _, v := range value {
					or = append(or, sq.Eq{key: v})
				}
				and = append(and, or)
			} else {
				and = append(and, sq.Eq{key: value[0]})
			}
		case "departure_time", "departure_date", "arrival_time", "arrival_date", "price":
			if len(value) > 1 {
				and = append(and, sq.And{sq.GtOrEq{key: value[1]}, sq.LtOrEq{key: value[0]}})
			} else {
				and = append(and, sq.Eq{key: value[0]})
			}
		default:
			and = append(and, sq.Eq{key: value[0]})
		}
		count = 0
	}

	req := flights.Where(and)

	request, _, err := req.ToSql()
	if err != nil {
		return []Flights{}, errors.New("ERROR: Bad request")
	}
	and = nil

	rows, err := DB.Query(request)
	if err != nil {
		return []Flights{}, err
	}

	flightsArray := make([]Flights, 0)

	for rows.Next() {
		var m Flights
		if err := rows.Scan(&m.ID, &m.departureCity, &m.departureTime, &m.departureDate, &m.arrivalCity, &m.arrivalDate, &m.arrivalTime, &m.price); err != nil {
			return []Flights{}, err
		}
		flightsArray = append(flightsArray, m)
	}
	return flightsArray, nil
}
