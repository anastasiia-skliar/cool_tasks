package models

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/satori/go.uuid"
	"reflect"
	"strings"
	"net/url"
	"time"
)
//Event is a representation of Event table in DB
type Event struct {
	ID       uuid.UUID
	Title    string
	Category string
	Town     string
	Date     time.Time
	Price    int
}
//Flight is a representation of Flight table in DB
type Flight struct {
	ID            uuid.UUID
	DepartureCity string
	DepartureTime time.Time
	DepartureDate time.Time
	ArrivalCity   string
	ArrivalTime   time.Time
	ArrivalDate   time.Time
	Price         int
}

var AddToTrip = func(dataID uuid.UUID, tripID uuid.UUID, dataSource interface{}) error {
	_, err := database.DB.Exec(generateQueryAdd(dataSource), dataID, tripID)
	return err
}
var GetFromTrip = func(tripID uuid.UUID, dataSource interface{}) (interface{}, error) {
	dataType := reflect.TypeOf(dataSource)
	rows, err := database.DB.Query(generateQueryGet(dataSource), tripID)
	if err != nil {
		return nil, err
	}
	switch dataType.Name() {
	case "Event":
		events := make([]Event, 0)
		for rows.Next() {
			var e Event
			if err := rows.Scan(&e.ID, &e.Title, &e.Category, &e.Town, &e.Date, &e.Price); err != nil {
				return []Event{}, err
			}
			events = append(events, e)
		}
		return events, nil
	case "Flight":
		flights := make([]Flight, 0)
		for rows.Next() {
			var f Flight
			if err := rows.Scan(&f.ID, &f.DepartureCity, &f.DepartureTime, &f.DepartureDate, &f.ArrivalCity, &f.ArrivalDate, &f.ArrivalTime, &f.Price); err != nil {
				return nil, err
			}
			flights = append(flights, f)
		}
		return flights, nil
	}

	return nil, nil
}

var GetData = func (params url.Values,dataSource interface{})(interface{}, error){
	dataType := reflect.TypeOf(dataSource)
	switch dataType.Name() {

	case "Event":
		stringArgs := []string{"title", "category", "town"}
		numberArgs := []string{"price", "date"}
		request, args, err := SQLGenerator("events", stringArgs, numberArgs, params)
		if err != nil {
			return nil, err
		}
		rows, err := database.DB.Query(request, args...)
		if err != nil {
			return []Event{}, err
		}
		events := make([]Event, 0)
		for rows.Next() {
			var e Event
			if err := rows.Scan(&e.ID, &e.Title, &e.Category, &e.Town, &e.Date, &e.Price); err != nil {
				return []Event{}, err
			}
			events = append(events, e)
		}
		return events, nil
	case "Flight":
		stringArgs := []string{"departure_city", "arrival_city"}
		numberArgs := []string{"price", "departure_time", "arrival_time", "departure_date", "arrival_date"}
		request, args, err := SQLGenerator("flights", stringArgs, numberArgs, params)
		if err != nil {
			return nil, err
		}
		rows, err := database.DB.Query(request, args...)
		if err != nil {
			return nil, err
		}

		flightsArray := make([]Flight, 0)

		for rows.Next() {
			var f Flight
			if err := rows.Scan(&f.ID, &f.DepartureCity, &f.DepartureTime, &f.DepartureDate, &f.ArrivalCity, &f.ArrivalDate, &f.ArrivalTime, &f.Price); err != nil {
				return nil, err
			}
			flightsArray = append(flightsArray, f)
		}
		return flightsArray, nil
	}
	return nil,nil
}

func generateQueryAdd(dataSource interface{}) string {
	dataType := reflect.TypeOf(dataSource)
	var query = "INSERT INTO trips_" + strings.ToLower(dataType.Name()) + "s" + " (" + strings.ToLower(dataType.Name()) + "_id, trip_id) VALUES ($1, $2)"
	return query
}

func generateQueryGet(dataSource interface{}) string {
	dataType := reflect.TypeOf(dataSource)
	name := strings.ToLower(dataType.Name())
	pluralName := name + "s"
	var query = "SELECT " + pluralName + ".* FROM " + pluralName + " INNER JOIN trips_" + pluralName + " ON " + pluralName + ".id=trips_" + pluralName + "." + name + "_id AND trips_" + pluralName + ".trip_id=$1"
	return query
}
