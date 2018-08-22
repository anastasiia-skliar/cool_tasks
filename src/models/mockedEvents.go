package models

import (
	"github.com/satori/go.uuid"
	"net/url"
)

//MockedGetEvents is mocked GetEvents func
func MockedGetEvents(events []Event, err error) {
	GetData = func(values url.Values, dataSource interface{}) (interface{}, error) {
		return events, err
	}
}

//MockedAddEventToTrip is mocked AddEventToTrip func
func MockedAddEventToTrip(err error) {
	AddToTrip = func(event_id uuid.UUID, trip_id uuid.UUID,dataSource interface{}) error {
		return err
	}
}

//MockedGetEventsByTrip is mocked GetEventsByTrip func
func MockedGetEventsByTrip(events []Event, err error) {
	GetFromTrip = func(trip_id uuid.UUID,dataSource interface{}) (interface{}, error) {
		return events, err
	}
}
