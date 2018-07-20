package models

import "github.com/satori/go.uuid"

//MockedGetTripsByTripID is mocked GetTripByTripID func
func MockedGetTripsByTripID(trip Trip) {
	GetTripByTripID = func(id uuid.UUID) (Trip, error) {

		return trip, nil
	}
}
