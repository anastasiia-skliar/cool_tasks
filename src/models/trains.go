package models

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/satori/go.uuid"
	"net/url"
	"time"
)

const (
	saveTrainToTrip  = "INSERT INTO trips_trains (trip_id, train_id) VALUES ($1, $2);"
	getTrainFromTrip = "SELECT trains.* FROM trains INNER JOIN trips_trains ON trips_trains.train_id = trains.id AND trips_trains.trip_id = $1;"
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
var GetTrains = func(params url.Values) ([]Train, error) {
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

//SaveTrain used for saving trains to Trip
var SaveTrain = func(tripsID, trainsID uuid.UUID) error {
	_, err := database.DB.Exec(saveTrainToTrip, tripsID, trainsID)

	return err
}

//GetTrainFromTrip used for getting trains from Trip
var GetTrainFromTrip = func(tripsID uuid.UUID) ([]Train, error) {
	rows, err := database.DB.Query(getTrainFromTrip, tripsID)
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
