package models

import (
	"database/sql"
	"fmt"
	. "github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/satori/go.uuid"
	sq "gopkg.in/Masterminds/squirrel.v1"
	"log"
	"net/url"
)

const (
	datalocation     = "restaurants"
	create           = "INSERT INTO %s (%s) VALUES (%s) RETURNING id"
	deleteTempl      = "DELETE FROM %s WHERE id = $1"
	saveToTripTempl  = "INSERT INTO trips_%s (trips_id, %s_id) VALUES ($1, $2)"
	getFromTripTempl = "SELECT * FROM %s INNER JOIN trips_%s ON trips_%s.%s_id = trains.id AND trips_%s.trips_id = $1"
)

var (
	deleteRequest   string
	getItemFromTrip string
)

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
	getItemFromTrip = getFromTripGenreator(datalocation)
}

var getFromTripGenreator = func(dest string) string {
	return fmt.Sprintf(getFromTripTempl, dest, dest, dest, dest, dest)
}

var recGen = func(params map[string][]string) (string, []interface{}){
	var cond sq.And
	var request string

	items := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).Select("*").From(datalocation)

	for k, v := range params {
		switch k {
		case "stars", "prices", "location", "name":
			if len(params[k]) > 1 {
				var multivars sq.Or
				for _, v2 := range params[k] {
					multivars = append(multivars, sq.Eq{k: v2})
				}
				cond = append(cond, multivars)
			} else {
				cond = append(cond, sq.Eq{k: v[0]})
			}
		case "departure_city", "arrival_city":
			cond = append(cond, sq.Eq{k: v[0]})

		}
	}
	var err error
	request, args, err := items.Where(cond).ToSql()
	if err != nil {
		log.Println(err)
	}
	if len(params) == 0 {
		request = fmt.Sprintf("SELECT * FROM %s",datalocation)
	}
	fmt.Println(request)
	return request, args
}

var parseResult = func(rows *sql.Rows) ([]Restaurant, error) {
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

//CreateTask used for creation task in DB
var AddRestaurant = func(item Restaurant) (Restaurant, error) {
	err := DB.QueryRow(create, item.Name, item.Location, item.Stars, item.Prices, item.Description).Scan(&item.ID)
	return item, err
}

//GetTask used for getting task from DB
var GetRestaurantsByID = func(id uuid.UUID) (Restaurant, error) {
	var item Restaurant
	params := make(map[string][]string)
	params["id"] = []string{id.String()}
	err := DB.QueryRow(recGen(params)).Scan(&item.ID, &item.Name, &item.Location, &item.Stars, &item.Prices, &item.Description)
	return item, err
}

//DeleteTask used for deleting task from DB
var DeleteRestaurantsFromDB = func(id uuid.UUID) error {
	_, err := DB.Exec(deleteRequest, id)
	return err
}

//GetTasks used for getting tasks from DB

var GetRestaurantsByQuery = func(query url.Values) ([]Restaurant, error) {
	sqlQuery, args:=recGen(query)
	rows, err := DB.Query(sqlQuery, args...)

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

var SaveRestaurantToTrip = func(tripID, restaurantID uuid.UUID, dataloc string) error {
	return saveToTrip(tripID, restaurantID, datalocation)
}

var saveToTrip = func(tripsID, itemID uuid.UUID, dataloc string) error {
	saveSQL := fmt.Sprintf(saveToTripTempl, dataloc, dataloc)
	_, err := DB.Exec(saveSQL, tripsID, itemID)

	return err
}

var GetRestaurantsFromTrip = func(tripsID uuid.UUID) ([]Restaurant, error) {
	rows, err := DB.Query(getTrainsFromTrip, tripsID)
	if err != nil {
		return []Restaurant{}, err
	}

	restaurants := make([]Restaurant, 0)
	restaurants, err = parseResult(rows)
	if err != nil {
		return []Restaurant{}, err
	}
	return restaurants, nil
}
