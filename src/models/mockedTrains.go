package models

import (
	"github.com/satori/go.uuid"
	"net/url"
)

func MockedGetTrains() {
	GetTrains = func(params url.Values) ([]Train, error) {
		return []Train{}, nil
	}
}

func MockedSaveTrain() {
	SaveTrain = func(tripsID, trainsID uuid.UUID) error {
		return nil
	}
}

func MockedGetFromTrip() {
	GetTrainFromTrip = func(tripsID uuid.UUID) ([]Train, error) {
		return []Train{}, nil
	}
}
