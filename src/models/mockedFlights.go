package models

import (
	"github.com/satori/go.uuid"
	"net/url"
)

//MockedAddToTrip is mocked AddToTrip func
func MockedAddToTrip(err error) {
	AddFlightToTrip = func(flightID, tripID uuid.UUID) error {
		return err
	}
}

//MockedGetByRequest is mocked GetByRequest func
func MockedGetByRequest(flights []Flight, err error) {
	GetFlights = func(params url.Values) ([]Flight, error) {
		return flights, err
	}
}

//MockedGetByTrip is mocked GetByTrip func
func MockedGetByTrip(flights []Flight, err error) {
	GetFlightsByTrip = func(tripID uuid.UUID) ([]Flight, error) {
		return flights, err
	}
}
