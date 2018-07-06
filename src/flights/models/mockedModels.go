package models

import (
	"github.com/satori/go.uuid"
	"time"
)

func MockedGetFlights(flights []Flights, err error) {
	GetFlights = func() ([]Flights, error) {
		return flights, err
	}
}

func MockedGetByCity(departureCity, arrrivalCity string, err error) {
	GetByCity = func(departureCity, arrivalCity string) ([]Flights, error) {
		return []Flights{}, err
	}
}

func MockedGetByDepartureTime(from, to time.Time, err error) {
	GetByDepartureTime = func(from, to time.Time) ([]Flights, error) {
		return []Flights{}, err
	}
}

func MockedGetByArrivalTime(from, to time.Time, err error) {
	GetByArrivalTime = func(from, to time.Time) ([]Flights, error) {
		return []Flights{}, err
	}
}

func MockedGetByPrice(from, to int, err error) {
	GetByPrice = func(from, to int) ([]Flights, error) {
		return []Flights{}, err
	}
}

func MockedAddToTrip(flightID, tripID uuid.UUID, err error) {
	AddToTrip = func(flightID, tripID uuid.UUID) (error) {
		return err
	}
}

func MockedGetByDate(departureDate, arrivalDate time.Time, err error) {
	GetByDate = func(departureDate, arrivalDate time.Time) ([]Flights, error) {
		return []Flights{}, err
	}
}

func MockedGetByTrip(tripID uuid.UUID, err error) {
	GetByTrip = func(tripID uuid.UUID) ([]Flights, error) {
		return []Flights{}, err
	}
}