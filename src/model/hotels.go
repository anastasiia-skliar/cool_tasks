package model

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/satori/go.uuid"
	"net/url"
)

const (
	addHotelToTrip = "INSERT INTO trips_hotels (trip_id, hotels_id) VALUES ($1, $2)"
	getHotelByTrip = "SELECT hotels.* FROM hotels INNER JOIN trips_hotels ON hotels.id=trips_hotels.hotels_id AND trips_hotels.trip_id=$1"
)

//Hotel is a representation of Hotel table in DB
type Hotel struct {
	ID        uuid.UUID
	Name      string
	Class     int
	Capacity  int
	RoomsLeft int
	Floors    int
	MaxPrice  int
	CityName  string
	Address   string
}

//AddHotelToTrip adds Hotel to Trip
var AddHotelToTrip = func(tripID, hotelID uuid.UUID) error {
	_, err := database.DB.Exec(addHotelToTrip, tripID, hotelID)
	return err
}

//GetHotelsByTrip gets Hotels from Trip by tripID
var GetHotelsByTrip = func(tripID uuid.UUID) ([]Hotel, error) {

	rows, err := database.DB.Query(getHotelByTrip, tripID)
	if err != nil {
		return nil, err
	}
	hotels := make([]Hotel, 0)
	for rows.Next() {
		var h Hotel
		if err := rows.Scan(&h.ID, &h.Name, &h.Class, &h.Capacity, &h.RoomsLeft, &h.Floors, &h.MaxPrice, &h.CityName, &h.Address); err != nil {
			return nil, err
		}
		hotels = append(hotels, h)
	}
	return hotels, nil
}

//GetHotels gets Hotels from Trip by incoming request
var GetHotels = func(params url.Values) ([]Hotel, error) {
	stringArgs := []string{"name", "city_name", "address"}
	numberArgs := []string{"class", "capacity", "rooms_left", "floors", "max_price"}
	request, args, err := SQLGenerator("hotels", stringArgs, numberArgs, params)
	if err != nil {
		return nil, err
	}

	rows, err := database.DB.Query(request, args...)
	if err != nil {
		return nil, err
	}

	hotels := make([]Hotel, 0)
	for rows.Next() {
		var h Hotel
		if err := rows.Scan(&h.ID, &h.Name, &h.Class, &h.Capacity, &h.RoomsLeft, &h.Floors, &h.MaxPrice, &h.CityName, &h.Address); err != nil {
			return nil, err
		}
		hotels = append(hotels, h)
	}
	return hotels, nil
}
