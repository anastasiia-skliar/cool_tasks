package models

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/satori/go.uuid"
	"time"
)

const (
	saveTrainToTrip   = "INSERT INTO trips_trains (trips_id, trains_id) VALUES ($1, $2)"
	getTrainsFromTrip = "SELECT * FROM trains INNER JOIN trips_trains ON trips_trains.trains_id = trains.id AND trips_trains.trips_id = $1"
)

//Task representation in DB
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

//GetTrains used for getting trains from DB
var GetTrains = func(query string) ([]Train, error) {
	rows, err := database.DB.Query(query)
	if err != nil {
		return []Train{}, err
	}

	trains := make([]Train, 0)
	for rows.Next() {
		var t Train
		if err := rows.Scan(&t.ID, &t.DepartureTime, &t.DepartureDate, &t.ArrivalTime, &t.ArrivalDate,
			&t.DepartureCity, &t.ArrivalCity, &t.TrainType, &t.CarType, &t.Price); err != nil {
			return []Train{}, err
		}
		trains = append(trains, t)
	}
	return trains, nil
}

//SaveTrain used for saving trains to Trip
var SaveTrain = func(tripsID, trainsID uuid.UUID) error {
	_, err := database.DB.Exec(saveTrainToTrip, tripsID, trainsID)

	return err
}

//GetTrainsFromTrip used for getting trains from Trip
var GetTrainsFromTrip = func(tripsID uuid.UUID) ([]Train, error) {
	rows, err := database.DB.Query(getTrainsFromTrip, tripsID)
	if err != nil {
		return []Train{}, err
	}

	trains := make([]Train, 0)
	for rows.Next() {
		var t Train
		if err := rows.Scan(&t.ID, &t.DepartureTime, &t.DepartureDate, &t.ArrivalTime, &t.ArrivalDate,
			&t.DepartureCity, &t.ArrivalCity, &t.TrainType, &t.CarType, &t.Price); err != nil {
			return []Train{}, err
		}
		trains = append(trains, t)
	}
	return trains, nil
}
