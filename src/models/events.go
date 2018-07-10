package models

import (
	"time"
	"github.com/satori/go.uuid"
	. "github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"net/url"
	"errors"
	"strings"
)

const (
	addEventToTrip = "INSERT INTO trips_flights (flight_id, trip_id) VALUES ($1, $2)"
	getEventByTrip = "SELECT * FROM flights INNER JOIN trips_evets ON events.id=trips_events.event_id AND trips_events.trip_id=$1"
)

type Event struct {
	ID        uuid.UUID
	Title     string
	Category  string
	Town 	  string
	Date      time.Time
	Price     int
}
var AddEventToTrip = func(eventID uuid.UUID, tripID uuid.UUID) (error) {
	_, err := DB.Exec(addEventToTrip, eventID, tripID)
	return err
}

var GetEventByTrip = func(eventID uuid.UUID) ([]Event, error) {
	rows, err := DB.Query(getEventByTrip, eventID)
	if err != nil {
		return []Event{}, err
	}

	events := make([]Event, 0)

	for rows.Next() {
		var p Event
		if err := rows.Scan(&p.ID, &p.Title, &p.Category, &p.Town, &p.Date, &p.Price); err != nil {
			return []Event{}, err
		}
		events = append(events, p)
	}
	return events, nil
}
var GetEventByRequest = func(params url.Values) ([]Event, error) {

	request := "SELECT * FROM events WHERE "
	count := 0
	validKeys := []string{"id","title","category","town","date","price"}
	for key, value := range params {
		for _, keys := range validKeys {
			if key == keys {
				count++
			}
		}
		if count == 0 {
			return []Event{}, errors.New("ERROR: Invalid request")

		}

		switch key {
		case "town","category":
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
		case "date","price":
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
		return []Event{}, err
	}

	events := make([]Event, 0)

	for rows.Next() {
		var p Event
		if err := rows.Scan(&p.ID, &p.Title, &p.Category, &p.Town, &p.Date, &p.Price); err != nil {
			return []Event{}, err

		}
		events = append(events, p)
	}
	return events, nil
}


