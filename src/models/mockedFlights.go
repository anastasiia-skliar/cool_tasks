package models

import (
	"github.com/satori/go.uuid"
	"net/url"
)

func MockedAddToTrip() {
	AddFlightToTrip = func(flightID, tripID uuid.UUID) (error) {
		return nil
	}
}

func MockedGetByRequest(){
	GetFlightsByRequest = func(params url.Values) ([]Flight, error) {
		return []Flight{}, nil
	}
}

func MockedGetByTrip(){
	GetFlightsByTrip = func(tripID uuid.UUID) ([]Flight, error) {
		return []Flight{}, nil
	}
}