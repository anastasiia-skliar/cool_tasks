package models

import (
	"errors"
	sq "github.com/Masterminds/squirrel"
	. "github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/satori/go.uuid"
	"net/url"
	"time"
)

const (
	addEventToTrip = "INSERT INTO trips_events (event_id, trip_id) VALUES ($1, $2)"
	getEventByTrip = "SELECT * FROM events INNER JOIN trips_events ON events.id=trips_events.event_id AND trips_events.trip_id=$1"
)

type Event struct {
	ID       uuid.UUID
	Title    string
	Category string
	Town     string
	Date     time.Time
	Price    int
}

var AddEventToTrip = func(eventID uuid.UUID, tripID uuid.UUID) error {
	_, err := DB.Exec(addEventToTrip, eventID, tripID)
	return err
}

var GetEventsByTrip = func(tripID uuid.UUID) ([]Event, error) {

	rows, err := DB.Query(getEventByTrip, tripID)
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

var GetEventsByRequest = func(params url.Values) ([]Event, error) {

	var and sq.And = nil
	events := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).Select("*").From("events")
	for key, value := range params {
		var or sq.Or = nil
		switch key {
		case "town", "title", "category":
			if len(value) > 1 {
				for _, v := range value {
					or = append(or, sq.Eq{key: v})
				}
				and = append(and, or)
			} else {
				and = append(and, sq.Eq{key: value[0]})
			}
		case "date", "price":
			if len(value) > 1 {
				and = append(and, sq.And{sq.GtOrEq{key: value[1]}, sq.LtOrEq{key: value[0]}})
			} else {
				and = append(and, sq.Eq{key: value[0]})
			}
		case "id":
			and = append(and, sq.Eq{key: value[0]})
		default:
			return []Event{}, errors.New("ERROR: Bad request")
		}
	}
	req := events.Where(and)
	request, _, err := req.ToSql()
	if err != nil {
		return []Event{}, errors.New("ERROR: Bad request")
	}
	rows, err := DB.Query(request)
	if err != nil {
		return []Event{}, err
	}
	req_events := make([]Event, 0)
	for rows.Next() {
		var e Event
		if err := rows.Scan(&e.ID, &e.Title, &e.Category, &e.Town, &e.Date, &e.Price); err != nil {
			return []Event{}, err
		}
		req_events = append(req_events, e)
	}
	return req_events, nil
}
