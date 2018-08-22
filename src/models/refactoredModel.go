package models

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/satori/go.uuid"
	"net/url"
	"reflect"
	"strings"
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

//Museum is a representation of Museum table in DB
type Museum struct {
	ID         uuid.UUID
	Name       string
	Location   string
	Price      int
	OpenedAt   time.Time
	ClosedAt   time.Time
	MuseumType string
	Info       string
}

//Train representation in DB
type Train struct {
	ID            uuid.UUID
	DepartureTime time.Time
	DepartureDate time.Time
	ArrivalTime   time.Time
	ArrivalDate   time.Time
	DepartureCity string
	ArrivalCity   string
	TrainType     string
	CarType       string
	Price         string
}

var AddToTrip = func(dataID uuid.UUID, tripID uuid.UUID, dataSource interface{}) error {
	_, err := database.DB.Exec(GenerateQueryAdd(dataSource), dataID, tripID)
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
	case "Museum":
		museums := make([]Museum, 0)

		for rows.Next() {
			var m Museum
			if err := rows.Scan(&m.ID, &m.Name, &m.Location, &m.Price, &m.OpenedAt, &m.ClosedAt, &m.MuseumType, &m.Info); err != nil {
				return nil, err
			}
			museums = append(museums, m)
		}
		return museums, nil
	case "Train":
		trains := make([]Train, 0)
		for rows.Next() {
			var t Train
			if err := rows.Scan(&t.ID, &t.DepartureTime, &t.DepartureDate, &t.ArrivalTime, &t.ArrivalDate,
				&t.DepartureCity, &t.ArrivalCity, &t.TrainType, &t.CarType, &t.Price); err != nil {
				return nil, err
			}
			trains = append(trains, t)
		}
		return trains, nil
	}
	return nil, nil
}

var GetData = func(params url.Values, dataSource interface{}) (interface{}, error) {
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
	case "Museum":
		stringArgs := []string{"name", "location", "museum_type"}
		numberArgs := []string{"price", "opened_at", "closed_at"}
		request, args, err := SQLGenerator("museums", stringArgs, numberArgs, params)
		if err != nil {
			return nil, err
		}
		rows, err := database.DB.Query(request, args...)
		if err != nil {
			return nil, err
		}

		museumsArray := make([]Museum, 0)

		for rows.Next() {
			var m Museum
			if err := rows.Scan(&m.ID, &m.Name, &m.Location, &m.Price, &m.OpenedAt, &m.ClosedAt, &m.MuseumType, &m.Info); err != nil {
				return nil, err
			}
			museumsArray = append(museumsArray, m)
		}
		return museumsArray, nil
	case "Train":
		stringArgs := []string{"departure_city", "arrival_city"}
		numberArgs := []string{"price", "departure_time", "arrival_time", "departure_date", "arrival_date"}
		request, args, err := SQLGenerator("trains", stringArgs, numberArgs, params)
		if err != nil {
			return nil, err
		}
		rows, err := database.DB.Query(request, args...)
		if err != nil {
			return nil, err
		}

		trains := make([]Train, 0)
		for rows.Next() {
			var t Train
			if err := rows.Scan(&t.ID, &t.DepartureTime, &t.DepartureDate, &t.ArrivalTime, &t.ArrivalDate,
				&t.DepartureCity, &t.ArrivalCity, &t.TrainType, &t.CarType, &t.Price); err != nil {
				return nil, err
			}
			trains = append(trains, t)
		}
		return trains, nil

	}
	return nil, nil
}

func GenerateQueryAdd(dataSource interface{}) string {
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
