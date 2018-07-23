package models

import (
	"github.com/satori/go.uuid"
	"net/url"
)

//MockedGetTrains is mocked GetTrains func
func MockedGetTrains() {
	GetTrains = func(params url.Values) ([]Train, error) {
		return []Train{}, nil
	}
}

//MockedSaveTrain is mocked AddTrainToTrip func
func MockedSaveTrain() {
	AddTrainToTrip = func(tripsID, trainsID uuid.UUID) error {
		return nil
	}
}

//MockedGetTrainsFromTrip is mocked GetTrainsFromTrip func
func MockedGetTrainsFromTrip() {
	GetTrainsFromTrip = func(tripsID uuid.UUID) ([]Train, error) {
		return []Train{}, nil
	}
}
