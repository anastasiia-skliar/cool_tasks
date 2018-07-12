package models

import (
	"github.com/satori/go.uuid"
	"net/url"
)

func MockedGetEvents() {
	GetEventsByRequest = func(values url.Values) ([]Event, error) {
		return []Event{}, nil
	}
}

func MockedAddEventToTrip() {
	AddEventToTrip = func(event_id uuid.UUID, trip_id uuid.UUID) error {
		return nil
	}
}

func MockedGetEventsByTrip() {
	GetEventsByTrip = func(trip_id uuid.UUID) ([]Event, error) {
		return []Event{}, nil
	}
}
