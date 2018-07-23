package models

import (
	sq "github.com/Masterminds/squirrel"
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
	MaxPrice  string
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

	var (
		cond    sq.And
		request string
		args    []interface{}
	)

	selectHotels := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).Select("*").From("hotels")

	for k, v := range params {
		switch k {
		case "id", "name", "class", "capacity", "room_left", "floors", "max_price", "city_name", "address":
			if len(params[k]) == 2 {
				cond = append(cond, sq.And{sq.GtOrEq{k: v[0]}, sq.LtOrEq{k: v[1]}})
			} else {
				cond = append(cond, sq.Eq{k: v[0]})
			}
		}
	}

	request, args, err := selectHotels.Where(cond).ToSql()
	if err != nil {
		return nil, err
	}

	if len(params) == 0 {
		request = "SELECT * FROM hotels;"
	}

	rows, err := database.DB.Query(request, args...)
	if err != nil {
		return []Hotel{}, err
	}

	hotels := make([]Hotel, 0)
	for rows.Next() {
		var h Hotel
		if err := rows.Scan(&h.ID, &h.Name, &h.Class, &h.Capacity, &h.RoomsLeft, &h.Floors, &h.MaxPrice, &h.CityName, &h.Address); err != nil {
			return []Hotel{}, err
		}
		hotels = append(hotels, h)
	}
	return hotels, nil
}
