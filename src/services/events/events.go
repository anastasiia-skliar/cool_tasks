//Package events implements event handlers
package events

import (
	"encoding/json"
	"github.com/Nastya-Kruglikova/cool_tasks/src/models"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/common"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"net/http"
)

type success struct {
	Status string `json:"message"`
}
type TripEvent struct {
	EventID string
	TripID  string
}

//AddEventToTripHandler is a handler for adding Event to Trip
func AddEventToTripHandler(w http.ResponseWriter, r *http.Request) {
	var newEvent TripEvent

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newEvent)
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Can't decode JSON POST Body", err)
		return
	}

	eventID, err := uuid.FromString(newEvent.EventID)
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong eventID (can't convert string to uuid)", err)
		return
	}

	tripID, err := uuid.FromString(newEvent.TripID)
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

//GetEventsByTripHandler is a handler for getting Events from Trip
func GetEventsByTripHandler(w http.ResponseWriter, r *http.Request) {
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

//GetEventsHandler is a handler for getting Events from Trip by request
func GetEventsHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	events, err := models.GetEvents(params)
	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't find events with such parameters", err)
		return
	}
	common.RenderJSON(w, r, events)
}
