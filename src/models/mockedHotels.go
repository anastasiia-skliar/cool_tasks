package models

import (
	"github.com/satori/go.uuid"
)

func GetHotelsMocked() {
	GetHotels = func(query string) ([]Hotel, error) {
		return []Hotel{}, nil
	}
}

func AddHotelMocked() {
	AddHotelToTrip = func(tripsID, hotelsID uuid.UUID) error {
		return nil
	}
}

func GetHotelByTripIdMocked() {
	GetHotelFromTrip = func(tripsID uuid.UUID) ([]Hotel, error) {
		return []Hotel{}, nil
	}
}
