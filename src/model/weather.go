package model

import (
	"encoding/json"
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/satori/go.uuid"
	"net/http"
	"time"
)

const (
	getCityTrain  = "select arrival_city from trains inner join trips_trains on trips_trains.train_id = trains.id and trips_trains.train_id =$1"
	getCityFlight = "select arrival_city from flights inner join trips_flights on trips_flights.flight_id = flights.id and trips_flights.flight_id =$1"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

type Forecast struct {
	Id      int
	Name    string
	Weather []Describe
	Main    Temp
}
type Describe struct {
	Id          int
	Main        string
	Description string
}
type Temp struct {
	Temp     float64
	Temp_min float64
	Temp_max float64
}

//GetWeatherByTrainId is used for getting weather in arrival city
var GetWeatherByTrainId = func(id uuid.UUID) (Forecast, error) {
	var city string
	var forecast Forecast
	err := database.DB.QueryRow(getCityTrain, id).Scan(&city)
	getJson("http://api.openweathermap.org/data/2.5/weather?q="+city+"&units=metric&appid=17c3332648cc87498de6efc362674d0d", &forecast)

	return forecast, err
}

//GetWeatherByFlightId is used for getting weather in arrival city
var GetWeatherByFlightId = func(id uuid.UUID) (Forecast, error) {
	var city string
	var forecast Forecast
	err := database.DB.QueryRow(getCityFlight, id).Scan(&city)
	getJson("http://api.openweathermap.org/data/2.5/weather?q="+city+"&units=metric&appid=17c3332648cc87498de6efc362674d0d", &forecast)

	return forecast, err
}
