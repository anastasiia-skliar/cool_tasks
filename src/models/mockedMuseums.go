package models

import (
	"github.com/satori/go.uuid"
	"net/url"
)

//MockedGetMuseums is mocked GetMuseums func
func MockedGetMuseums(museum []Museum, err error) {
	GetData = func(values url.Values, dataSource interface{}) (interface{}, error) {
		return museum, err
	}
}

//MockedAddMuseum is mocked AddMuseum func
func MockedAddMuseum(err error) {
	AddToTrip = func(museum_id uuid.UUID, trip_id uuid.UUID, dataSource interface{}) error {
		return err
	}
}

//MockedGetMuseumsByTrip is mocked GetMuseumsByTrip func
func MockedGetMuseumsByTrip(museum []Museum, err error) {
	GetFromTrip = func(trip_id uuid.UUID, dataSource interface{}) (interface{}, error) {
		return museum, err
	}
}
