package models

import (
	"github.com/satori/go.uuid"
	. "github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"time"
	"strings"
	"net/url"
	"errors"
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

var GetByRequest = func(params url.Values) ([]Flights, error) {

	request := "SELECT * FROM flights WHERE "
	count := 0
	validKeys := []string{"id", "departure_city", "departure_time", "departure_date", "arrival_city", "arrival_time", "arrival_date", "price"}
	for key, value := range params {
		for _, keys := range validKeys {
			if key == keys {
				count++
			}
		}
		if count == 0 {
			return []Flights{}, errors.New("ERROR: Invalid request")

		}

		switch key {
		case "departure_city", "arrival_city":
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
		case "departure_time", "departure_date", "arrival_time", "arrival_date", "price":
			if len(value) > 1 {
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
	request += ";"

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