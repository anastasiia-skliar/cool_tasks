package models

import (
	. "github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/satori/go.uuid"
	"net/url"
)

const (
	addHotelToTrip = "INSERT INTO trips_hotels (trip_id, hotels_id) VALUES ($1, $2)"
	getHotelByTrip = "SELECT hotels.* FROM hotels INNER JOIN trips_hotels ON hotels.id=trips_hotels.hotels_id AND trips_hotels.trip_id=$1"
)

type Hotel struct {
	ID         uuid.UUID
	NAME       string
	CLASS      int
	CAPACITY   int
	ROOMS_LEFT int
	FLOORS     int
	MAX_PRICE  string
	CITY_NAME  string
	ADDRESS    string
}

var AddHotelToTrip = func(tripID, hotelID uuid.UUID) error {
	_, err := DB.Exec(addHotelToTrip, tripID, hotelID)
	return err
}

var GetHotelsByTrip = func(tripID uuid.UUID) ([]Hotel, error) {

	rows, err := DB.Query(getHotelByTrip, tripID)
	if err != nil {
		return nil, err
	}
	hotels := make([]Hotel, 0)
	for rows.Next() {
		var h Hotel
		if err := rows.Scan(&h.ID, &h.NAME, &h.CLASS, &h.CAPACITY, &h.ROOMS_LEFT, &h.FLOORS, &h.MAX_PRICE, &h.CITY_NAME, &h.ADDRESS); err != nil {
			return nil, err
		}
		hotels = append(hotels, h)
	}
	return hotels, nil
}

var GetHotelsByRequest = func(params url.Values) ([]Hotel, error) {
	stringArgs := []string{"name", "city_name", "address"}
	numberArgs := []string{"class", "capacity", "rooms_left", "floors", "max_price"}
	request, args, err := SQLGenerator("hotels", stringArgs, numberArgs, params)
	if err != nil {
		return nil, err
	}

	rows, err := DB.Query(request, args...)
	if err != nil {
		return nil, err
	}

	hotels := make([]Hotel, 0)
	for rows.Next() {
		var h Hotel
		if err := rows.Scan(&h.ID, &h.NAME, &h.CLASS, &h.CAPACITY, &h.ROOMS_LEFT, &h.FLOORS, &h.MAX_PRICE, &h.CITY_NAME, &h.ADDRESS); err != nil {
			return nil, err
		}
		hotels = append(hotels, h)
	}
	return hotels, nil
}
