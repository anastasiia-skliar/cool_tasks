package model

import (
	"github.com/satori/go.uuid"
	"net/url"
)

//MockedGetEvents is mocked GetEvents func
func MockedGetEvents(events []Event, err error) {
	GetEvents = func(values url.Values) ([]Event, error) {
		return events, err
	}
}

//MockedAddEventToTrip is mocked AddEventToTrip func
func MockedAddEventToTrip(err error) {
	AddEventToTrip = func(event_id uuid.UUID, trip_id uuid.UUID) error {
		return err
	}
}

//MockedGetEventsByTrip is mocked GetEventsByTrip func
func MockedGetEventsByTrip(events []Event, err error) {
	GetEventsByTrip = func(trip_id uuid.UUID) ([]Event, error) {
		return events, err
	}
}
