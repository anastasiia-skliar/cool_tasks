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
	addFlightToTrip = "INSERT INTO trips_flights (flight_id, trip_id) VALUES ($1, $2)"
	getFlightByTrip = "SELECT * FROM flights INNER JOIN trips_flights ON flights.id=trips_flights.flight_id AND trips_flights.trip_id=$1"
)

type Flight struct {
	ID            uuid.UUID
	departureCity string
	departureTime time.Time
	departureDate time.Time
	arrivalCity   string
	arrivalTime   time.Time
	arrivalDate   time.Time
	price         int
}

var AddFlightToTrip = func(flightID uuid.UUID, tripID uuid.UUID) (error) {
	_, err := DB.Exec(addFlightToTrip, flightID, tripID)
	return err
}

var GetFlightsByTrip = func(tripID uuid.UUID) ([]Flight, error) {
	rows, err := DB.Query(getFlightByTrip, tripID)
	if err != nil {
		return []Flight{}, err
	}

	flights := make([]Flight, 0)

	for rows.Next() {
		var f Flight
		if err := rows.Scan(&f.ID, &f.departureCity, &f.departureTime, &f.departureDate, &f.arrivalCity, &f.arrivalDate, &f.arrivalTime, &f.price); err != nil {
			return []Flight{}, err
		}
		flights = append(flights, f)
	}
	return flights, nil
}

var GetFlightsByRequest = func(params url.Values) ([]Flight, error) {

	var and sq.And
	var or sq.Or
	flights := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).Select("*").From("flights")
	for key, value := range params {
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
		case "id":
			and = append(and, sq.Eq{key: value[0]})
		default:
			return []Flight{}, errors.New("ERROR: Bad request")
		}
	}

	req := flights.Where(and)

	request, _, err := req.ToSql()
	if err != nil {
		return []Flight{}, errors.New("ERROR: Bad request")
	}
	and = nil //make and nil after request

	rows, err := DB.Query(request)
	if err != nil {
		return []Flight{}, err
	}

	flightsArray := make([]Flight, 0)

	for rows.Next() {
		var f Flight
		if err := rows.Scan(&f.ID, &f.departureCity, &f.departureTime, &f.departureDate, &f.arrivalCity, &f.arrivalDate, &f.arrivalTime, &f.price); err != nil {
			return []Flight{}, err
		}
		flightsArray = append(flightsArray, f)
	}
	return flightsArray, nil
}
