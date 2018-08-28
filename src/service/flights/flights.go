//Package flights implements flight handlers
package flights

import (
	"encoding/json"
	"github.com/Nastya-Kruglikova/cool_tasks/src/model"
	"github.com/Nastya-Kruglikova/cool_tasks/src/service/common"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"net/http"
)

type success struct {
	Status string `json:"message"`
}

type TripFlight struct {
	FlightID string `json:"flight_id"`
	TripID   string `json:"trip_id"`
}

//AddFlightToTripHandler is a handler for adding Flight to Trip
func AddFlightToTripHandler(w http.ResponseWriter, r *http.Request) {
	var newFlight TripFlight

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newFlight)
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Can't decode JSON POST Body", err)
		return
	}

	flightID, err := uuid.FromString(newFlight.FlightID)
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong flight ID (can't convert string to uuid)", err)
		return
	}

	tripID, err := uuid.FromString(newFlight.TripID)
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong trip ID (can't convert string to uuid)", err)
		return
	}

	err = model.AddFlightToTrip(flightID, tripID)
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Can't add new flight to trip", err)
		return
	}
	common.RenderJSON(w, r, success{Status: "201 Created"})
}

//GetFlightsByTripHandler is a handler for getting Flight from Trip
func GetFlightsByTripHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	tripID, err := uuid.FromString(params["id"])
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong trip ID (can't convert string to uuid)", err)
		return
	}

	flights, err := model.GetFlightsByTrip(tripID)
	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't get flights by trip ID", err)
		return
	}
	common.RenderJSON(w, r, flights)
}

//GetFlightsHandler is a handler for getting Flight from Trip by request
func GetFlightsHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	flights, err := model.GetFlights(params)
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Can't find flights with such parameters", err)
		return
	}
	common.RenderJSON(w, r, flights)
}
