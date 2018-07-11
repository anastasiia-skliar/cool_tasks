package models

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/satori/go.uuid"
)

const (
	createTrip = "INSERT INTO trips (user_id) VALUES ($1) RETURNING trip_id"
	getTrips   = "SELECT trains.*, museums.* FROM trips LEFT JOIN trips_trains ON trips_trains.trip_id = trips.trip_id LEFT JOIN trains ON trips_trains.train_id = trains.idLEFT JOIN trips_museums ON trips_museums.trip_id = trips.trip_id LEFT JOIN museums ON trips_museums.museum_id = museums.id WHERE trips.user_id = 'a46889d4-839f-11e8-8d89-c01885c5bc39';"
)

type Trip struct {
	TripID uuid.UUID
	UserID uuid.UUID
}

var CreateTrip = func(trip Trip) (uuid.UUID, error) {
	var id uuid.UUID
	err := database.DB.QueryRow(createTrip, trip.UserID).Scan(&id)

	return id, err
}

var GetTripsByID  = func(id uuid.UUID) ([]Trip, error) {

	return nil, nil
}
