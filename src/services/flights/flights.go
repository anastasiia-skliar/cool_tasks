package flights

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

func AddToTripHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Can't parse POST Body", err)
		return
	}

	flightID, err := uuid.FromString(r.Form.Get("flight_id"))

	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong flight ID (can't convert string to uuid)", err)
		return
	}

	tripID, err := uuid.FromString(r.Form.Get("trip_id"))

	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong trip ID (can't convert string to uuid)", err)
		return
	}

	err = models.AddToTrip(flightID, tripID)

	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Can't add new flight to trip", err)
		return
	}

	common.RenderJSON(w, r, success{Status: "201 Created"})
}

func GetByTripHandler(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	tripID, err := uuid.FromString(params["id"])

	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong trip ID (can't convert string to uuid)", err)
		return
	}

	flights, err := models.GetByTrip(tripID)

	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't get flights by trip ID", err)
		return
	}

	common.RenderJSON(w, r, flights)
}

func GetByRequestHandler(w http.ResponseWriter, r *http.Request) {

	params := r.URL.Query()

	flights, err := models.GetByRequest(params)

	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't find flights with such parameters", err)
		return
	}

	common.RenderJSON(w, r, flights)
}
