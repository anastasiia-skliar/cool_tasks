//Package model contains functions interacting with DB
package model

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/satori/go.uuid"
	"net/url"
	"time"
)

const (
	addEventToTrip = "INSERT INTO trips_events (event_id, trip_id) VALUES ($1, $2)"
	getEventByTrip = "SELECT events.* FROM events INNER JOIN trips_events ON events.id=trips_events.event_id AND trips_events.trip_id=$1"
)

//Event is a representation of Event table in DB
type Event struct {
	ID       uuid.UUID
	Title    string
	Category string
	Town     string
	Date     time.Time
	Price    int
}

//AddEventToTrip adds Events to Trip
var AddEventToTrip = func(eventID uuid.UUID, tripID uuid.UUID) error {
	_, err := database.DB.Exec(addEventToTrip, eventID, tripID)
	return err
}

//GetEventsByTrip gets Events from Trip by tripID
var GetEventsByTrip = func(tripID uuid.UUID) ([]Event, error) {

	rows, err := database.DB.Query(getEventByTrip, tripID)
	if err != nil {
		return nil, err
	}
	events := make([]Event, 0)
	for rows.Next() {
		var e Event
		if err := rows.Scan(&e.ID, &e.Title, &e.Category, &e.Town, &e.Date, &e.Price); err != nil {
			return []Event{}, err
		}
		events = append(events, e)
	}
	return events, nil
}

//GetEvents gets Events by incoming request
var GetEvents = func(params url.Values) ([]Event, error) {

	stringArgs := []string{"title", "category", "town"}
	numberArgs := []string{"price", "date"}
	request, args, err := SQLGenerator("events", stringArgs, numberArgs, params)
	if err != nil {
		return nil, err
	}
	rows, err := database.DB.Query(request, args...)
	if err != nil {
		return []Event{}, err
	}
	events := make([]Event, 0)
	for rows.Next() {
		var e Event
		if err := rows.Scan(&e.ID, &e.Title, &e.Category, &e.Town, &e.Date, &e.Price); err != nil {
			return []Event{}, err
		}
		events = append(events, e)
	}
	return events, nil
}
