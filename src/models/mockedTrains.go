package models

import (
	"github.com/satori/go.uuid"
	"net/url"
)

//MockedGetTrains is mocked GetTrains func
func MockedGetTrains(train []Train, err error) {
	GetData = func(params url.Values, dataSource interface{}) (interface{}, error) {
		return train, err
	}
}

//MockedSaveTrain is mocked AddTrainToTrip func
func MockedSaveTrain(err error) {
	AddToTrip = func(tripsID, trainsID uuid.UUID, dataSource interface{}) error {
		return err
	}
}

//MockedGetTrainsFromTrip is mocked GetTrainsFromTrip func
func MockedGetTrainsFromTrip(train []Train, err error) {
	GetFromTrip = func(tripsID uuid.UUID, dataSource interface{}) (interface{}, error) {
		return train, err
	}
}
