package models

import (
	"github.com/satori/go.uuid"
	"net/url"
)

func GetHotelsMocked() {
	GetHotelsByRequest = func(params url.Values) ([]Hotel, error) {
		return []Hotel{}, nil
	}
}

func AddHotelMocked() {
	AddHotelToTrip = func(tripsID, hotelsID uuid.UUID) error {
		return nil
	}
}

func GetHotelByTripIdMocked() {
	GetHotelsByTrip = func(tripsID uuid.UUID) ([]Hotel, error) {
		return []Hotel{}, nil
	}
}
