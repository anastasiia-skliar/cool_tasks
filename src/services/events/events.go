package events

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

//AddToTripHandler is a handler for adding Event to Trip
func AddToTripHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Can't parse POST Body", err)
		return
	}

	eventID, err := uuid.FromString(r.Form.Get("event_id"))
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong eventID (can't convert string to uuid)", err)
		return
	}

	tripID, err := uuid.FromString(r.Form.Get("trip_id"))
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong tripID (can't convert string to uuid)", err)
		return
	}

	err = models.AddEventToTrip(eventID, tripID)
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Can't add new event to trip", err)
		return
	}
	common.RenderJSON(w, r, success{Status: "201 Created"})
}

//GetByTripHandler is a handler for getting Events from Trip
func GetByTripHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	tripID, err := uuid.FromString(params["id"])
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong tripID (can't convert string to uuid)", err)
		return
	}

	events, err := models.GetEventsByTrip(tripID)
	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't get events by tripID", err)
		return
	}
	common.RenderJSON(w, r, events)
}

//GetByRequestHandler is a handler for getting Events from Trip by request
func GetByRequestHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	events, err := models.GetEventsByRequest(params)
	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't find events with such parameters", err)
		return
	}
	common.RenderJSON(w, r, events)
}
