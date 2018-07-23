package models

import "github.com/satori/go.uuid"

//MockedGetTripsByTripID is mocked GetTrip func
func MockedGetTripsByTripID(trip Trip) {
	GetTrip = func(id uuid.UUID) (Trip, error) {

		return trip, nil
	}
}
