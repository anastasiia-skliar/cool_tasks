package models

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/satori/go.uuid"
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

const (
	addHotelToTrip   = "INSERT INTO trips_hotels (trips_id, hotels_id) VALUES ($1, $2)"
	getHotelFromTrip = "SELECT * FROM hotels INNER JOIN trips_hotels ON trips_hotels.hotels_id = hotels.id AND trips_hotels.hotels_id = $1"
)

//get hotel from a db
var GetHotels = func(query string) ([]Hotel, error) {
	rows, err := database.DB.Query(query)
	if err != nil {
		return []Hotel{}, err
	}
	hotels := make([]Hotel, 0)
	for rows.Next() {
		var h Hotel
		if err := rows.Scan(&h.ID, &h.NAME, &h.CLASS, &h.CAPACITY, &h.ROOMS_LEFT, &h.FLOORS, &h.MAX_PRICE, &h.ADDRESS); err != nil {
			return []Hotel{}, err
		}
		hotels = append(hotels, h)
	}
	return hotels, nil
}

var AddHotelToTrip = func(tripsID, hotelsID uuid.UUID) error {
	_, err := database.DB.Exec(addHotelToTrip, tripsID, hotelsID)
	return err
}
var GetHotelFromTrip = func(tripsID uuid.UUID) ([]Hotel, error) {
	rows, err := database.DB.Query(getHotelFromTrip, tripsID)
	if err != nil {
		return []Hotel{}, err
	}
	hotels := make([]Hotel, 0)
	for rows.Next() {
		var h Hotel
		if err := rows.Scan(&h.ID, &h.NAME, &h.CLASS, &h.CAPACITY, &h.ROOMS_LEFT,
			&h.FLOORS, &h.MAX_PRICE, &h.ADDRESS); err != nil {
			return []Hotel{}, err
		}
		hotels = append(hotels, h)
	}
	return hotels, nil
}
