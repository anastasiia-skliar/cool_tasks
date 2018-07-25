package trips

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/models"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/common"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"net/http"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/auth"
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
	itemOwner, err:= models.GetUserByID(result.UserID)
	if auth.CheckPermission(r, "owner", itemOwner.Login)==false{
		common.SendError(w, r, http.StatusForbidden, "Wrong user role", nil)
		return
	}
	common.RenderJSON(w, r, result)
}

func GetTripIDByUserID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := uuid.FromString(params["id"])
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong userID", err)
		return
	}
	itemOwner, err:= models.GetUserByID(userID)
	if auth.CheckPermission(r, "owner", itemOwner.Login)==false{
		common.SendError(w, r, http.StatusForbidden, "Wrong user role", nil)
		return
	}
	result, err := models.GetTripIDByUserID(userID)

	common.RenderJSON(w, r, result)
}
