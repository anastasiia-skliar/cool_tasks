//Package museums implements museum handlers
package museums

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/model"
	"github.com/Nastya-Kruglikova/cool_tasks/src/service/common"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"net/http"
)

type successCreate struct {
	Status string `json:"status"`
}

//AddMuseumToTripHandler is a handler for adding Museums to Trips
func AddMuseumToTripHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Can't parse POST Body", err)
		return
	}
	museumID, err := uuid.FromString(r.Form.Get("museum_id"))
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Converting ID from POST Body", err)
		return
	}
	tripID, err := uuid.FromString(r.Form.Get("trip_id"))
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Converting ID from POST Body", err)
		return
	}

	musErr := model.AddToTrip(museumID, tripID, model.Museum{})
	if musErr != nil {
		common.SendBadRequest(w, r, "ERROR: Cant ADD Museum", err)
		return
	}
	common.RenderJSON(w, r, successCreate{Status: "201 Created"})
}

//GetMuseumsByTripHandler is a handler for getting Museums from Trips
func GetMuseumsByTripHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tripID, err := uuid.FromString(params["id"])
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Converting ID from URL", err)
		return
	}
	museums, err := model.GetFromTrip(tripID, model.Museum{})
	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't find museums in such trip", err)
		return
	}
	common.RenderJSON(w, r, museums)
}

//GetMuseumsHandler is a handler for getting Museums from Trips by request
func GetMuseumsHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	museums, err := model.GetFromTripWithParams(params, model.Museum{})
	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't find museums with such parameters", err)
		return
	}
	common.RenderJSON(w, r, museums)
}
