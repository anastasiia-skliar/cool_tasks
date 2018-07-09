package trains

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/models"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/common"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"net/http"
)

type successAdd struct {
	Status string       `json:"message"`
	Result models.Train `json:"result"`
}

func GetTrains(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	trains, err := models.GetTrains(params)
	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't find any trains", err)
		return
	}

	common.RenderJSON(w, r, trains)
}

func SaveTrain(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Can't parse POST Body", err)
		return
	}

	trainID, err := uuid.FromString(r.Form.Get("train_id"))
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong train ID (can't convert string to uuid)", err)
		return
	}

	tripID, err := uuid.FromString(r.Form.Get("trip_id"))
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong trip ID (can't convert string to uuid)", err)
		return
	}

	err = models.SaveTrain(trainID, tripID)
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Can't add new train to trip", err)
		return
	}

	common.RenderJSON(w, r, successAdd{Status: "201 Created"})
}

func GetFromTrip(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	tripID, err := uuid.FromString(params["id"])
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong trip ID (can't convert string to uuid)", err)
		return
	}

	trains, err := models.GetFromTrip(tripID)
	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't get trains by trip ID", err)
		return
	}

	common.RenderJSON(w, r, trains)
}
