package model

import (
	"database/sql"
	"github.com/satori/go.uuid"
	"net/url"
)

//MockedGetRestaurants is mocked GetMuseums func
func MockedGetRestaurants(res []Restaurant, err error) {
	GetRestaurants = func(values url.Values) ([]Restaurant, error) {
		return res, err
	}
}

//MockedAddMuseum is mocked AddMuseum func
func MockedAddRestaurants(err error) {
	AddRestaurantToTrip = func(restaurant_id uuid.UUID, trip_id uuid.UUID) error {
		return err
	}
}

//MockedGetMuseumsByTrip is mocked GetMuseumsByTrip func
func MockedGetRestaurantsByTrip(res []Restaurant, err error) {
	GetRestaurantsFromTrip = func(trip_id uuid.UUID) ([]Restaurant, error) {
		return res, err
	}
}

func MockedParseResult(res []Restaurant, err error) {
	parseResult = func(rows *sql.Rows) ([]Restaurant, error) {
		return res, err
	}
}
