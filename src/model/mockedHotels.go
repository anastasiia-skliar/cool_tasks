package model

import (
	"github.com/satori/go.uuid"
	"net/url"
)

//GetHotelsMocked is mocked GetHotels func
func GetHotelsMocked(hotels []Hotel, err error) {
	GetHotels = func(params url.Values) ([]Hotel, error) {
		return hotels, err
	}
}

//AddHotelMocked is mocked AddHotel func
func AddHotelMocked(err error) {
	AddHotelToTrip = func(tripsID, hotelsID uuid.UUID) error {
		return err
	}
}

//GetHotelByTripIDMocked is mocked GetHotelByTripID func
func GetHotelByTripIDMocked(hotels []Hotel, err error) {
	GetHotelsByTrip = func(tripsID uuid.UUID) ([]Hotel, error) {
		return hotels, err
	}
}
