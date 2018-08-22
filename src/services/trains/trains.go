//Package trains implements train handlers
package trains

import (
	"net/http"

	"github.com/Nastya-Kruglikova/cool_tasks/src/models"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/common"

	"fmt"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
)

type successAdd struct {
	Status string `json:"message"`
}

//AddTrainToTripHandler is a handler for saving Train to Trip
func AddTrainToTripHandler(w http.ResponseWriter, r *http.Request) {
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

	err = models.AddToTrip(trainID, tripID, models.Train{})
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Can't add new train to trip", err)
		fmt.Println(err)
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

	trains, err := models.GetFromTrip(tripID, models.Train{})
	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't get trains by trip ID", err)
		return
	}

	common.RenderJSON(w, r, trains)
}

//GetTrainsHandler is a handler for getting Train from Trip by request
func GetTrainsHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	trains, err := models.GetData(params, models.Train{})
	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't find any trains", err)
		return
	}

	common.RenderJSON(w, r, trains)
}
