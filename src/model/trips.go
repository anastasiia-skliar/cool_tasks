package model

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/satori/go.uuid"
	"log"
)

const (
	createTrip        = "INSERT INTO trips (user_id) VALUES ($1) RETURNING trip_id;"
	getTrip           = "SELECT * FROM trips WHERE trip_id = $1;"
	getTripIDByUserID = "SELECT trips.trip_id FROM trips WHERE trips.user_id = $1;"
)

//Trip is a representation of Event Trip in DB
type Trip struct {
	TripID      uuid.UUID
	UserID      uuid.UUID
	Events      []Event
	Flights     []Flight
	Museums     []Museum
	Restaurants []Restaurant
	Hotels      []Hotel
	Trains      []Train
	TotalSum    int
}

//CreateTrip creates Trip and saves it to DB
var CreateTrip = func(trip Trip) (uuid.UUID, error) {
	var id uuid.UUID
	err := database.DB.QueryRow(createTrip, trip.UserID).Scan(&id)

	return id, err
}

//GetTrip gets Trips from DB by tripID
var GetTrip = func(id uuid.UUID) (Trip, error) {

	var (
		trip Trip
		err  error
	)

	dbErr := database.DB.QueryRow(getTrip, id).Scan(&trip.TripID, &trip.UserID)
	if dbErr != nil {
		log.Println(dbErr)
		return Trip{}, dbErr
	}

	trip.Events, err = GetEventsByTrip(id)
	if err != nil {
		log.Println(err)
		return Trip{}, err
	}
	for _, item := range trip.Events {
		trip.TotalSum += item.Price
	}

	trip.Flights, err = GetFlightsByTrip(id)
	if err != nil {
		log.Println(err)
		return Trip{}, err
	}
	for _, item := range trip.Flights {
		trip.TotalSum += item.Price
	}

	trip.Museums, err = GetMuseumsByTrip(id)
	if err != nil {
		log.Println(err)
		return Trip{}, err
	}
	for _, item := range trip.Museums {
		trip.TotalSum += item.Price
	}

	trip.Hotels, err = GetHotelsByTrip(id)
	if err != nil {
		log.Println(err)
		return Trip{}, err
	}
	for _, item := range trip.Hotels {
		trip.TotalSum += item.MaxPrice
	}

	trip.Trains, err = GetTrainsFromTrip(id)
	if err != nil {
		log.Println(err)
		return Trip{}, err
	}
	for _, item := range trip.Trains {
		trip.TotalSum += item.Price
	}

	trip.Restaurants, err = GetRestaurantsFromTrip(id)
	if err != nil {
		log.Println(err)
		return Trip{}, err
	}
	for _, item := range trip.Restaurants {
		trip.TotalSum += item.Price
	}

	return trip, nil
}

//GetTripIDsByUserID gets Trips from DB by userID
var GetTripIDsByUserID = func(id uuid.UUID) ([]uuid.UUID, error) {

	var (
		tripIDs []uuid.UUID
		tripID  uuid.UUID
	)

	rows, err := database.DB.Query(getTripIDByUserID, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err := rows.Scan(&tripID); err != nil {
			return nil, err
		}
		tripIDs = append(tripIDs, tripID)
	}

	return tripIDs, nil
}
