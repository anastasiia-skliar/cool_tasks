package models

import "github.com/satori/go.uuid"

func MockedGetTripsByTripID(trip Trip) {
	GetTripsByTripID = func(id uuid.UUID) (Trip, error) {

		return trip, nil
	}
}
