package models

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/satori/go.uuid"
	"net/url"
	"time"
)

const (
	saveTrainToTrip  = "INSERT INTO trips_trains (trip_id, train_id) VALUES ($1, $2)"
	getTrainFromTrip = "SELECT * FROM trains INNER JOIN trips_trains ON trips_trains.train_id = trains.id AND trips_trains.trip_id = $1"
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

	var (
		cond    sq.And
		request string
	)

	selectTrains := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).Select("*").From("trains")

	for k, v := range params {
		switch k {
		case "id", "departure_time", "departure_date", "arrival_time", "arrival_date", "price":
			if len(params[k]) == 2 {
				cond = append(cond, sq.And{sq.GtOrEq{k: v[0]}, sq.LtOrEq{k: v[1]}})
			} else {
				cond = append(cond, sq.Eq{k: v[0]})
			}
		case "departure_city", "arrival_city":
			cond = append(cond, sq.Eq{k: v[0]})
		}
	}

	request, _, _ = selectTrains.Where(cond).ToSql()

	if len(params) == 0 {
		request = "SELECT * FROM trains;"
	}

	rows, err := database.DB.Query(request)
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

//GetTrainFromTrip used for getting trains from Trip
var GetTrainFromTrip = func(tripsID uuid.UUID) ([]Train, error) {
	rows, err := database.DB.Query(getTrainFromTrip, tripsID)
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
