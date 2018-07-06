package museums

import (
	"net/http"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/common"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
)

type successCreate struct {
	Status string    `json:"status"`
}

func GetMuseumsHandler(w http.ResponseWriter, r *http.Request) {
	museums, err := GetMuseums()
	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't get users", err)
		return
	}
	common.RenderJSON(w, r, museums)
}

func GetMuseumByCityHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	museums, err := GetMuseumsByCity(params["city"])
	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't find museums in such city", err)
		return
	}
	common.RenderJSON(w, r, museums)
}

func AddMuseumToTripHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Can't parse POST Body", err)
		return
	}
	museum_id, err := uuid.FromString(r.Form.Get("museum"))
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Converting ID from POST Body", err)
		return
	}
	trip_id, err := uuid.FromString(r.Form.Get("trip"))
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Converting ID from POST Body", err)
		return
	}
	err = AddMuseumToTrip(museum_id, trip_id)
	common.RenderJSON(w, r, successCreate{Status: "201 Created"})
}

func GetMuseumByTripHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tripID, err := uuid.FromString(params["id"])
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Converting ID from URL", err)
		return
	}
	museums, err := GetMuseumsByTrip(tripID)
	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't find museums in such trip", err)
		return
	}
	common.RenderJSON(w, r, museums)
}
