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
func AddEventToTripHandler(w http.ResponseWriter, r *http.Request) {

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

err = models.AddToTrip(eventID, tripID)

if err != nil {
common.SendBadRequest(w, r, "ERROR: Can't add new event to trip", err)
return
}

common.RenderJSON(w, r, success{Status: "201 Created"})
}

func GetEventByTripHandler(w http.ResponseWriter, r *http.Request) {

params := mux.Vars(r)

tripID, err := uuid.FromString(params["id"])

if err != nil {
common.SendBadRequest(w, r, "ERROR: Wrong tripID (can't convert string to uuid)", err)
return
}

	events, err := models.GetByTrip(tripID)

if err != nil {
common.SendNotFound(w, r, "ERROR: Can't get events by tripID", err)
return
}

common.RenderJSON(w, r, events)
}

func GetByRequestHandler(w http.ResponseWriter, r *http.Request) {

params := r.URL.Query()

	events, err := models.GetEventsByRequest(params)

if err != nil {
common.SendNotFound(w, r, "ERROR: Can't find events with such parameters", err)
return
}

common.RenderJSON(w, r, events)
}
