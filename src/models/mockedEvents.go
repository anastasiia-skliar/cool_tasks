package models

import (
	"github.com/satori/go.uuid"
	"net/url"
)

//MockedGetEvents is mocked GetEvents func
func MockedGetEvents() {
	GetEventsByRequest = func(values url.Values) ([]Event, error) {
		return []Event{}, nil
	}
}

//MockedAddEventToTrip is mocked AddEventToTrip func
func MockedAddEventToTrip() {
	AddEventToTrip = func(event_id uuid.UUID, trip_id uuid.UUID) error {
		return nil
	}
}

//MockedGetEventsByTrip is mocked GetEventsByTrip func
func MockedGetEventsByTrip() {
	GetEventsByTrip = func(trip_id uuid.UUID) ([]Event, error) {
		return []Event{}, nil
	}
}
