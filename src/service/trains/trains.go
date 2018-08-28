//Package trains implements train handlers
package trains

import (
	"net/http"

	"github.com/Nastya-Kruglikova/cool_tasks/src/model"
	"github.com/Nastya-Kruglikova/cool_tasks/src/service/common"

	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
)

type successAdd struct {
	Status string `json:"message"`
}

type TripTrain struct {
	TrainID string `json:"train_id"`
	TripID  string `json:"trip_id"`
}

//AddTrainToTripHandler is a handler for saving Train to Trip
func AddTrainToTripHandler(w http.ResponseWriter, r *http.Request) {
	var newTrain TripTrain

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newTrain)
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Can't decode JSON POST Body", err)
		return
	}

	trainID, err := uuid.FromString(newTrain.TrainID)
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong train ID (can't convert string to uuid)", err)
		return
	}

	tripID, err := uuid.FromString(newTrain.TripID)
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong trip ID (can't convert string to uuid)", err)
		return
	}

	err = model.AddToTrip(tripID, trainID,model.Train{})
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Can't add new train to trip", err)
		return
	}

	common.RenderJSON(w, r, successAdd{Status: "201 Created"})
}

//GetTrainsFromTrip is a handler for getting Trains from Trip
func GetTrainsFromTripHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	tripID, err := uuid.FromString(params["id"])
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong trip ID (can't convert string to uuid)", err)
		return
	}

	trains, err := model.GetFromTrip(tripID,model.Train{})
	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't get trains by trip ID", err)
		return
	}

	common.RenderJSON(w, r, trains)
}

//GetTrainsHandler is a handler for getting Train from Trip by request
func GetTrainsHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	trains, err := model.GetFromTripWithParams(params,model.Train{})
	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't find any trains", err)
		return
	}

	common.RenderJSON(w, r, trains)
}
