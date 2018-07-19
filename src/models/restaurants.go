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

//SaveRest saves Restaurant to Trip
var SaveRest = func(tripsID, restaurantsID uuid.UUID) error {
	_, err := database.DB.Exec(saveRestToTrip, tripsID, restaurantsID)

	return err
}

//GetRestByID gets Restaurants from Trip by tripID
var GetRestByID = func(id uuid.UUID) (Restaurant, error) {
	var item Restaurant
	err := database.DB.QueryRow(recGen("id"), id).Scan(&item.ID, &item.Name, &item.Location, &item.Stars, &item.Prices, &item.Description)
	return item, err
}

//DeleteRestFromDB deletes Restaurant from DB
var DeleteRestFromDB = func(id uuid.UUID) error {
	_, err := database.DB.Exec(deleteRequest, id)
	return err
}

//GetRestByQuery gets Restaurants from Trip by incoming query
var GetRestByQuery = func(query url.Values) ([]Restaurant, error) {

	paramNames := make([]string, 0)
	paramVals := make([]string, 0)
	for key, value := range query {
		if len(value) > 0 {
			for _, v := range value {
				paramNames = append(paramNames, key)
				paramVals = append(paramVals, v)
			}
		} else {
			continue
		}

	}

	s := make([]interface{}, len(paramVals))
	for i, v := range paramVals {
		s[i] = v
	}

	rows, err := database.DB.Query(recGen(paramNames...), s...)

	if err != nil {
		fmt.Println(err)
		return []Restaurant{}, err
	}
	res, err := parseResult(rows)
	if err != nil {
		fmt.Println(err)
		return []Restaurant{}, err
	}
	return res, nil
}

//GetRestFromTrip gets Restaurants from Trip by tripID
var GetRestFromTrip = func(tripsID uuid.UUID) ([]Restaurant, error) {
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
