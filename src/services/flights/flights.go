//Package flights implements flight handlers
package flights

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/models"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/common"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"net/http"
)

type success struct {
	Status string `json:"message"`
}

//AddFlightToTripHandler is a handler for adding Flight to Trip
func AddFlightToTripHandler(w http.ResponseWriter, r *http.Request) {
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

	err = models.AddFlightToTrip(flightID, tripID)
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

	flights, err := models.GetFlightsByTrip(tripID)
	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't get flights by trip ID", err)
		return
	}
	common.RenderJSON(w, r, flights)
}

//GetFlightsHandler is a handler for getting Flight from Trip by request
func GetFlightsHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	flights, err := models.GetFlights(params)
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Can't find flights with such parameters", err)
		return
	}
	common.RenderJSON(w, r, flights)
}
