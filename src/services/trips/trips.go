//Package events implements event handlers
package trips

import (
	"net/http"

	"github.com/Nastya-Kruglikova/cool_tasks/src/models"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/common"

	"github.com/satori/go.uuid"
	"github.com/gorilla/mux"
)

type successCreate struct {
	Status string    `json:"status"`
	ID     uuid.UUID `json:"id"`
}

//AddTripHandler is a handler for creating Trips
func AddTripHandler(w http.ResponseWriter, r *http.Request) {
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

	id, err := models.AddTrip(trip)
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Can't add this trip", err)
		return
	}
	common.RenderJSON(w, r, successCreate{Status: "201 Created", ID: id})
}

//GetTrip is a handler for getting Trip from DB bu tripID
func GetTripHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	tripID, err := uuid.FromString(params["id"])
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong tripID", err)
		return
	}

	result, err := models.GetTrip(tripID)
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Can't get this trip", err)
		return
	}

	common.RenderJSON(w, r, result)
}

//GetTripIDsByUserIDHandler is a handler for getting tripID from DB by userID
func GetTripIDsByUserIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := uuid.FromString(params["id"])
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong userID", err)
		return
	}

	result, err := models.GetTripIDsByUserID(userID)
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Can't get this trip", err)
		return
	}

	common.RenderJSON(w, r, result)
}
