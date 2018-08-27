package model

import (
	"database/sql"
	"fmt"
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/satori/go.uuid"
	"net/url"
)

const (
	datalocation    = "restaurants"
	saveRestToTrip  = "INSERT INTO trips_restaurants (trip_id, restaurant_id) VALUES ($1, $2);"
	deleteTempl     = "DELETE FROM %s WHERE id = $1"
	getRestFromTrip = "SELECT restaurants.* FROM restaurants INNER JOIN trips_restaurants ON trips_restaurants.restaurant_id = restaurants.id AND trips_restaurants.trip_id = $1;"
)

var deleteRequest string

//Restaurant representation in DB
type Restaurant struct {
	ID          uuid.UUID
	Name        string
	Location    string
	Stars       int
	Price       int
	Description string
}

func init() {
	deleteRequest = fmt.Sprintf(deleteTempl, datalocation)
}

func parseResult(rows *sql.Rows) ([]Restaurant, error) {
	res := make([]Restaurant, 0)

	for rows.Next() {
		var item Restaurant
		if err := rows.Scan(&item.ID, &item.Name, &item.Location, &item.Stars, &item.Price, &item.Description); err != nil {
			return []Restaurant{}, err
		}
		res = append(res, item)
	}
	return res, nil
}

//AddRestaurantToTrip saves Restaurant to Trip
var AddRestaurantToTrip = func(tripsID, restaurantsID uuid.UUID) error {
	_, err := database.DB.Exec(saveRestToTrip, tripsID, restaurantsID)

	return err
}

//DeleteRestaurant deletes Restaurant from DB
var DeleteRestaurant = func(id uuid.UUID) error {
	_, err := database.DB.Exec(deleteRequest, id)
	return err
}

//GetRestaurants gets Restaurants from Trip by incoming query
var GetRestaurants = func(params url.Values) ([]Restaurant, error) {
	stringArgs := []string{"id", "name", "location"}
	numberArgs := []string{"stars", "prices"}
	request, args, err := SQLGenerator(datalocation, stringArgs, numberArgs, params)
	if err != nil {
		return nil, err
	}
	rows, err := database.DB.Query(request, args...)
	if err != nil {
		return nil, err
	}
	return parseResult(rows)
}

//GetRestaurantsFromTrip gets Restaurants from Trip by tripID
var GetRestaurantsFromTrip = func(tripsID uuid.UUID) ([]Restaurant, error) {
	rows, err := database.DB.Query(getRestFromTrip, tripsID)
	if err != nil {
		return nil, err
	}
	return parseResult(rows)
}
