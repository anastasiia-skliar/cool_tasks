package models

import (
	"database/sql"
	"fmt"
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/satori/go.uuid"
	"net/url"
)

const (
	datalocation    = "restaurants"
	getter          = "SELECT * FROM %s"
	saveRestToTrip  = "INSERT INTO trips_restaurants (trip_id, restaurant_id) VALUES ($1, $2);"
	getByParameter  = "WHERE %s = $1"
	addParam        = " AND %s = $%d"
	addOr           = " OR %s = $%d"
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
	Prices      int
	Description string
}

func init() {
	deleteRequest = fmt.Sprintf(deleteTempl, datalocation)
}

func recGen(params ...string) string {
	base := fmt.Sprintf(getter, datalocation)
	if len(params) < 1 {
		return base
	}
	paramsCounter := 0
	request := fmt.Sprintf(base+" "+getByParameter, params[paramsCounter])
	paramsCounter++
	for ; paramsCounter < len(params); paramsCounter++ {
		if params[paramsCounter] != params[paramsCounter-1] {
			request += fmt.Sprintf(addParam, params[paramsCounter], paramsCounter+1)
		} else {
			request += fmt.Sprintf(addOr, params[paramsCounter], paramsCounter+1)
		}

	}
	fmt.Println(request)
	return request
}

func parseResult(rows *sql.Rows) ([]Restaurant, error) {
	res := make([]Restaurant, 0)

	for rows.Next() {
		var item Restaurant
		if err := rows.Scan(&item.ID, &item.Name, &item.Location, &item.Stars, &item.Prices, &item.Description); err != nil {
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

//GetRestaurant gets Restaurants from Trip by tripID
var GetRestaurant = func(id uuid.UUID) (Restaurant, error) {
	var item Restaurant
	err := database.DB.QueryRow(recGen("id"), id).Scan(&item.ID, &item.Name, &item.Location, &item.Stars, &item.Prices, &item.Description)
	return item, err
}

//DeleteRestaurant deletes Restaurant from DB
var DeleteRestaurant = func(id uuid.UUID) error {
	_, err := database.DB.Exec(deleteRequest, id)
	return err
}

//GetRestaurants gets Restaurants from Trip by incoming query
var GetRestaurants = func(params url.Values) ([]Restaurant, error) {
	stringArgs := []string{"name", "location"}
	numberArgs := []string{"stars", "prices"}
	request, args, err := SQLGenerator("restaurants", stringArgs, numberArgs, params)
	if err != nil {
		return nil, err
	}

	rows, err := database.DB.Query(request, args...)
	if err != nil {
		return nil, err
	}

	restaurant := make([]Restaurant, 0)
	for rows.Next() {
		var r Restaurant
		if err := rows.Scan(&r.ID, &r.Name, &r.Location, &r.Stars, &r.Prices, &r.Description); err != nil {
			return nil, err
		}
		restaurant = append(restaurant, r)
	}
	return restaurant, nil
}

//GetRestaurantsFromTrip gets Restaurants from Trip by tripID
var GetRestaurantsFromTrip = func(tripsID uuid.UUID) ([]Restaurant, error) {
	rows, err := database.DB.Query(getRestFromTrip, tripsID)
	if err != nil {
		return nil, err
	}

	restaurants := make([]Restaurant, 0)
	for rows.Next() {
		var r Restaurant
		if err := rows.Scan(&r.ID, &r.Name, &r.Location, &r.Stars, &r.Prices, &r.Description); err != nil {
			return nil, err
		}
		restaurants = append(restaurants, r)
	}
	return restaurants, nil
}
