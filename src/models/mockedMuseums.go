package models

import (
	"github.com/satori/go.uuid"
	"net/url"
)

//MockedGetMuseums is mocked GetMuseums func
func MockedGetMuseums() {
	GetMuseumsByRequest = func(values url.Values) ([]Museum, error) {
		return []Museum{}, nil
	}
}

//MockedAddMuseum is mocked AddMuseum func
func MockedAddMuseum() {
	AddMuseumToTrip = func(museum_id uuid.UUID, trip_id uuid.UUID) error {
		return nil
	}
}

//MockedGetMuseumsByTrip is mocked GetMuseumsByTrip func
func MockedGetMuseumsByTrip() {
	GetMuseumsByTrip = func(trip_id uuid.UUID) ([]Museum, error) {
		return []Museum{}, nil
	}
}
