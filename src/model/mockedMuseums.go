package model

import (
	"github.com/satori/go.uuid"
	"net/url"
)

//MockedGetMuseums is mocked GetMuseums func
func MockedGetMuseums(museum []Museum, err error) {
	GetMuseums = func(values url.Values) ([]Museum, error) {
		return museum, err
	}
}

//MockedAddMuseum is mocked AddMuseum func
func MockedAddMuseum(err error) {
	AddMuseumToTrip = func(museum_id uuid.UUID, trip_id uuid.UUID) error {
		return err
	}
}

//MockedGetMuseumsByTrip is mocked GetMuseumsByTrip func
func MockedGetMuseumsByTrip(museum []Museum, err error) {
	GetMuseumsByTrip = func(trip_id uuid.UUID) ([]Museum, error) {
		return museum, err
	}
}
