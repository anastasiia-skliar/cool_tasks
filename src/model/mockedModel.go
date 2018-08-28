package model

import (
	"github.com/satori/go.uuid"
	"net/url"
)

func MockedGetData(obj interface{}, err error) {
	GetFromTripWithParams = func(values url.Values, dataSource interface{}) (interface{}, error) {
		return dataSource, err
	}
}

func MockedAddToTrip(err error) {
	AddToTrip = func(id uuid.UUID, trip_id uuid.UUID, dataSource interface{}) error {
		return err
	}
}

//MockedGetByTrip is mocked GetByTrip func
func MockedGetByTrip(obj interface{}, err error) {
	GetFromTrip = func(tripID uuid.UUID, dataSource interface{}) (interface{}, error) {
		return dataSource, err
	}
}
