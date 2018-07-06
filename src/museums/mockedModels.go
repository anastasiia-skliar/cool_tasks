package museums

import (
	"github.com/satori/go.uuid"
)

func MockedGetMuseums() {
	GetMuseumsByRequest = func(string) ([]Museum, error) {
		return []Museum{}, nil
	}
}

func MockedAddMuseum() {
	AddMuseumToTrip = func(museum_id uuid.UUID, trip_id uuid.UUID) error {
		return nil
	}
}

func MockedGetMuseumsByTrip() {
	GetMuseumsByTrip = func(trip_id uuid.UUID) ([]Museum, error) {
		return []Museum{}, nil
	}
}
