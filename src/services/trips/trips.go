package trips

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/models"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/common"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"net/http"
)

type successCreate struct {
	Status string    `json:"status"`
	ID     uuid.UUID `json:"id"`
}

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

func GetTripsByTripID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	tripID, err := uuid.FromString(params["id"])
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong tripID", err)
		return
	}
	result, err := models.GetTripsByTripID(tripID)

	common.RenderJSON(w, r, result)
}

func GetTripIDByUserID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := uuid.FromString(params["id"])
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong userID", err)
		return
	}
	result, err := models.GetTripIDByUserID(userID)

	common.RenderJSON(w, r, result)
}
