package models

import (
	"github.com/satori/go.uuid"
	"net/url"
)

//MockedAddToTrip is mocked AddToTrip func
func MockedAddToTrip(err error) {
	AddToTrip = func(flightID, tripID uuid.UUID, dataSource interface{}) error {
		return err
	}
}

//MockedGetByRequest is mocked GetByRequest func
func MockedGetByRequest(flights []Flight, err error) {
	GetData = func(params url.Values, dataSource interface{}) (interface{}, error) {
		return flights, err
	}
}

//MockedGetByTrip is mocked GetByTrip func
func MockedGetByTrip(flights []Flight, err error) {
	GetFromTrip = func(tripID uuid.UUID, dataSource interface{}) (interface{}, error) {
		return flights, err
	}
}
