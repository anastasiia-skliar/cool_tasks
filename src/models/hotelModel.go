package models

import (
	sq "github.com/Masterminds/squirrel"
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

	request, args, _ = selectHotels.Where(cond).ToSql()

	if len(params) == 0 {
		request = "SELECT * FROM hotels;"
	}

	rows, err := DB.Query(request, args...)
	if err != nil {
		return []Hotel{}, err
	}

	hotels := make([]Hotel, 0)
	for rows.Next() {
		var h Hotel
		if err := rows.Scan(&h.ID, &h.NAME, &h.CLASS, &h.CAPACITY, &h.ROOMS_LEFT, &h.FLOORS, &h.MAX_PRICE, &h.CITY_NAME, &h.ADDRESS); err != nil {
			return []Hotel{}, err
		}
		hotels = append(hotels, h)
	}
	return hotels, nil
}
