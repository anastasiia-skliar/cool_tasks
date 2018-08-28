//Package trips implements event handlers
package trips

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/model"
	"github.com/Nastya-Kruglikova/cool_tasks/src/service/auth"
	"github.com/Nastya-Kruglikova/cool_tasks/src/service/common"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"net/http"
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

	var trip model.Trip
	trip.UserID, err = uuid.FromString(r.Form.Get("user_id"))
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong userID", err)
		return
	}

	id, err := model.AddTrip(trip)
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Can't add this trip", err)
		return
	}
	common.RenderJSON(w, r, successCreate{Status: "201 Created", ID: id})
}

//GetTrip is a handler for getting Trip from DB by tripID
func GetTripHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	tripID, err := uuid.FromString(params["id"])
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong tripID", err)
		return
	}

	result, err := model.GetTrip(tripID)
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Can't get this trip", err)
		return
	}

	itemOwner, err := model.GetUserByID(result.UserID)
	sessionID, err := auth.GetSessionIDFromRequest(w, r)
	if err != nil {
		return
	}
	if auth.CheckPermission(sessionID, auth.Owner, itemOwner.Login) == false {
		common.SendError(w, r, http.StatusForbidden, auth.NotOwnerResponse, nil)
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
	itemOwner, err := model.GetUserByID(userID)
	sessionID, err := auth.GetSessionIDFromRequest(w, r)
	if err != nil {
		return
	}
	if auth.CheckPermission(sessionID, auth.Owner, itemOwner.Login) == false {
		common.SendError(w, r, http.StatusForbidden, auth.NotOwnerResponse, nil)
		return
	}

	result, err := model.GetTripIDsByUserID(userID)
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Can't get this trip", err)
		return
	}

	common.RenderJSON(w, r, result)
}
