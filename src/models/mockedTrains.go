package models

import (
	"github.com/satori/go.uuid"
)

func MockedGetTrains() {
	GetTrains = func(query string) ([]Train, error) {
		return []Train{}, nil
	}
}

func MockedSaveTrain() {
	SaveTrain = func(tripsID, trainsID uuid.UUID) error {
		return nil
	}
}

func MockedGetFromTrip() {
	GetFromTrip = func(tripsID uuid.UUID) ([]Train, error) {
		return []Train{}, nil
	}
}
