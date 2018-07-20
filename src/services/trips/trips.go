package trips

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/models"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/common"
	"github.com/satori/go.uuid"
	"net/http"
	"github.com/gorilla/mux"
)

type successCreate struct {
	Status string    `json:"status"`
	ID     uuid.UUID `json:"id"`
}

//CreateTrip is a handler for creating Trips
func CreateTrip(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Can't parse POST Body", err)
		return
	}

	var trip models.Trip
	trip.UserID, err = uuid.FromString(r.Form.Get("user_id"))
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong userID", err)
		return
	}

	id, err := models.CreateTrip(trip)
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Can't add this trip", err)
		return
	}
	common.RenderJSON(w, r, successCreate{Status: "201 Created", ID: id})
}

//GetTripByTripID is a handler for getting Trip from DB bu tripID
func GetTripsByTripID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	tripID, err := uuid.FromString(params["id"])
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong tripID", err)
		return
	}

	result, err := models.GetTripByTripID(tripID)
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Can't get this trip", err)
		return
	}

	common.RenderJSON(w, r, result)
}

//GetTripIDByUserID is a handler for getting tripID from DB by userID
func GetTripIDByUserID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := uuid.FromString(params["id"])
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong userID", err)
		return
	}

	result, err := models.GetTripIDByUserID(userID)
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Can't get this trip", err)
		return
	}

	common.RenderJSON(w, r, result)
}
