package models

import (
	"github.com/satori/go.uuid"
)

func MockedAddToTrip() {
	AddToTrip = func(flightID, tripID uuid.UUID) (error) {
		return nil
	}
}

func MockedGetByRequest(){
	GetByRequest = func(request string) ([]Flights, error) {
		return []Flights{}, nil
	}
}

func MockedGetByTrip(){
	GetByTrip = func(tripID uuid.UUID) ([]Flights, error) {
		return []Flights{}, nil
	}
}