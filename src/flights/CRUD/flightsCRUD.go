package CRUD

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/flights/models"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/common"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"net/http"
	"time"
	"strconv"
)

type success struct {
	Status string `json:"message"`
}

func GetFlights(w http.ResponseWriter, r *http.Request) {

	flights, err := models.GetFlights()

	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't get flights", err)
		return
	}

	common.RenderJSON(w, r, flights)
}

func GetByCity(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	departureCity := params["departure_city"]
	arrivalCity := params["arrival_city"]

	flights, err := models.GetByCity(departureCity, arrivalCity)

	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't get flights by city", err)
		return
	}

	common.RenderJSON(w, r, flights)
}

func GetByDepartureTime(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	departureTimeFrom, err := time.Parse(time.UnixDate, params["departure_time_from"])
	departureTimeTo, err := time.Parse(time.UnixDate, params["departure_time_to"])

	flights, err := models.GetByArrivalTime(departureTimeFrom, departureTimeTo)

	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't get flights by departure time", err)
		return
	}

	common.RenderJSON(w, r, flights)
}

func GetByArrivalTime(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	arrivalTimeFrom, err := time.Parse(time.UnixDate, params["arrival_time_from"])
	arrivalTimeTo, err := time.Parse(time.UnixDate, params["arrival_time_to"])

	flights, err := models.GetByArrivalTime(arrivalTimeFrom, arrivalTimeTo)

	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't get flights by arrival time", err)
		return
	}

	common.RenderJSON(w, r, flights)
}

func GetByPrice(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	priceFrom, err := strconv.Atoi(params["price_from"])

	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Converting price from string to int", err)
		return
	}

	priceTo, err := strconv.Atoi(params["price_to"])

	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Converting price from string to int", err)
		return
	}

	flights, err := models.GetByPrice(priceFrom, priceTo)

	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't get flights by arrival time", err)
		return
	}

	common.RenderJSON(w, r, flights)
}

func AddToTrip(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Can't parse POST Body", err)
		return
	}

	flightID, err := uuid.FromString(r.Form.Get("flight_id"))

	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong flight ID (can't convert string to uuid)", err)
		return
	}

	tripID, err := uuid.FromString(r.Form.Get("trip_id"))

	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong trip ID (can't convert string to uuid)", err)
		return
	}

	err = models.AddToTrip(flightID,tripID)

	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Can't add new flight to trip", err)
		return
	}

	common.RenderJSON(w, r, success{Status: "201 Created"})
}

func GetByDate(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	departureDate, err := time.Parse(time.UnixDate, params["departure_date"])
	arrivalDate, err := time.Parse(time.UnixDate, params["arrival_date"])

	flights, err := models.GetByArrivalTime(departureDate, arrivalDate)

	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't get flights by date", err)
		return
	}

	common.RenderJSON(w, r, flights)
}

func GetByTrip(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	tripID, err := uuid.FromString(params["trip_id"])

	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong trip ID (can't convert string to uuid)", err)
		return
	}

	flights, err := models.GetByTrip(tripID)

	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't get flights by trip ID", err)
		return
	}

	common.RenderJSON(w, r, flights)
}