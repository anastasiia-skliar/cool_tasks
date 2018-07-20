package models

import (
	"github.com/satori/go.uuid"
	"net/url"
)

//GetHotelsMocked is mocked GetHotels func
func GetHotelsMocked() {
	GetHotelsByRequest = func(params url.Values) ([]Hotel, error) {
		return []Hotel{}, nil
	}
}

//AddHotelMocked is mocked AddHotel func
func AddHotelMocked() {
	AddHotelToTrip = func(tripsID, hotelsID uuid.UUID) error {
		return nil
	}
}

//GetHotelByTripIDMocked is mocked GetHotelByTripID func
func GetHotelByTripIDMocked() {
	GetHotelsByTrip = func(tripsID uuid.UUID) ([]Hotel, error) {
		return []Hotel{}, nil
	}
}
