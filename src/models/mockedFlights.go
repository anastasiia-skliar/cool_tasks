package models

import (
	"github.com/satori/go.uuid"
	"net/url"
)

func MockedAddToTrip() {
	AddToTrip = func(flightID, tripID uuid.UUID) (error) {
		return nil
	}
}

func MockedGetByRequest(){
	GetByRequest = func(params url.Values) ([]Flights, error) {
		return []Flights{}, nil
	}
}

func MockedGetByTrip(){
	GetByTrip = func(tripID uuid.UUID) ([]Flights, error) {
		return []Flights{}, nil
	}
}