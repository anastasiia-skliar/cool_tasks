package model

import (
	"github.com/satori/go.uuid"
)

//MockedGetTripsByTripID is mocked GetTrip func
func MockedGetTripsByTripID(trip Trip) Trip {
	GetTrip = func(id uuid.UUID) (Trip, error) {
		trip := Trip{TripID: id}
		return trip, nil
	}
	return trip
}
