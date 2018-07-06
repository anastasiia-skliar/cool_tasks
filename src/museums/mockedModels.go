package museums

import (
	"github.com/satori/go.uuid"
)

func MockedGetMuseums() {
	GetMuseums = func() ([]Museum, error) {
		return []Museum{}, nil
	}
}

func MockedGetMuseumByCity() {
	MuseumId, _ := uuid.FromString("00000000-0000-0000-0000-000000000001")
	GetMuseumsByCity = func(city string) ([]Museum, error) {
		return []Museum{
			{
				MuseumId,
				"Louvre",
				"Paris",
				1111,
				1,
				2,
				"History",
				"Cool",
			},
		}, nil
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
