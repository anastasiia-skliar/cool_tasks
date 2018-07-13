package models

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/satori/go.uuid"
)

const (
	createTrip       = "INSERT INTO trips (user_id) VALUES ($1) RETURNING trip_id;"
	getTripsByUserID = "SELECT trips.trip_id FROM trips WHERE trips.user_id = $1;"
)

type Trip struct {
	TripID  uuid.UUID
	UserID  uuid.UUID
	Events  []Event
	Flights []Flight
	Museums []Museum
	Hotels  []Hotel
	Trains  []Train
}

var CreateTrip = func(trip Trip) (uuid.UUID, error) {
	var id uuid.UUID
	err := database.DB.QueryRow(createTrip, trip.UserID).Scan(&id)

	return id, err
}

var GetTripsByTripID = func(id uuid.UUID) (Trip, error) {

	var (
		trip    Trip
		events  []Event
		flights []Flight
		museums []Museum
		hotels  []Hotel
		trains  []Train
	)

	events, _ = GetEventsByTrip(id)
	flights, _ = GetFlightsByTrip(id)
	museums, _ = GetMuseumsByTrip(id)
	hotels, _ = GetHotelsByTrip(id)
	trains, _ = GetTrainFromTrip(id)

	trip.Events = events
	trip.Flights = flights
	trip.Museums = museums
	trip.Hotels = hotels
	trip.Trains = trains

	return trip, nil
}

var GetTripIDByUserID = func(id uuid.UUID) ([]uuid.UUID, error) {

	var (
		tripIDs []uuid.UUID
		tripID  uuid.UUID
	)

	rows, err := database.DB.Query(getTripsByUserID, id)
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
