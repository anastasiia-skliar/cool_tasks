package models

import (
	"github.com/satori/go.uuid"
	"net/url"
)

//MockedGetTrains is mocked GetTrains func
func MockedGetTrains(train []Train, err error) {
	GetTrains = func(params url.Values) ([]Train, error) {
		return train, err
	}
}

//MockedSaveTrain is mocked AddTrainToTrip func
func MockedSaveTrain(err error) {
	AddTrainToTrip = func(tripsID, trainsID uuid.UUID) error {
		return err
	}
}

//MockedGetTrainsFromTrip is mocked GetTrainsFromTrip func
func MockedGetTrainsFromTrip(train []Train, err error) {
	GetTrainsFromTrip = func(tripsID uuid.UUID) ([]Train, error) {
		return train, err
	}
}
