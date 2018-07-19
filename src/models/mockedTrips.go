package models

import "github.com/satori/go.uuid"

//MockedGetTripsByTripID is mocked GetTripsByTripID func
func MockedGetTripsByTripID(trip Trip) {
	GetTripsByTripID = func(id uuid.UUID) (Trip, error) {

		return trip, nil
	}
}
