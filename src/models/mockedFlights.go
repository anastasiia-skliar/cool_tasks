package models

import (
	"github.com/satori/go.uuid"
	"net/url"
)

//MockedAddToTrip is mocked AddToTrip func
func MockedAddToTrip() {
	AddFlightToTrip = func(flightID, tripID uuid.UUID) error {
		return nil
	}
}

//MockedGetByRequest is mocked GetByRequest func
func MockedGetByRequest() {
	GetFlightsByRequest = func(params url.Values) ([]Flight, error) {
		return []Flight{}, nil
	}
}

//MockedGetByTrip is mocked GetByTrip func
func MockedGetByTrip() {
	GetFlightsByTrip = func(tripID uuid.UUID) ([]Flight, error) {
		return []Flight{}, nil
	}
}
