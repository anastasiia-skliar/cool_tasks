package models

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/satori/go.uuid"
)

const (
	create = "INSERT INTO trips (user_id) VALUES ($1) RETURNING trip_id"
)

type Trip struct {
	TripID uuid.UUID
	UserID uuid.UUID
}

var CreateTrip = func(trip Trip) (uuid.UUID, error) {
	var id uuid.UUID
	err := database.DB.QueryRow(create, trip.UserID).Scan(&id)

	return id, err
}
