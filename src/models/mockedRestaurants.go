package models

import (
	"net/url"
)

//MockedGetRestaurants is mocked GetRestaurants func
func MockedGetRestaurants(restaurant []Restaurant, err error) {
	GetRestaurants = func(params url.Values) ([]Restaurant, error) {
		return restaurant, err
	}
}
