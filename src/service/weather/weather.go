//Package weather implements weather handlers
package weather

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/model"
	"github.com/Nastya-Kruglikova/cool_tasks/src/service/common"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"net/http"
)

//GetWeatherByTrainIDHandler is a handler for getting weather in arrival city in case using the train in trip
func GetWeatherByTrainIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	trainID, err := uuid.FromString(params["id"])
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong trainID", err)
		return
	}

	result, err := model.GetWeatherByTrainID(trainID)
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: This trainID is not connected with your trip", err)
		return
	}

	common.RenderJSON(w, r, result)
}

//GetWeatherByTrainIDHandler is a handler for getting weather in arrival city in case using the train in trip
func GetWeatherByFlightIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	flightID, err := uuid.FromString(params["id"])
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong flightID", err)
		return
	}

	result, err := model.GetWeatherByFlightID(flightID)
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: This flightID is not connected with your trip", err)
		return
	}

	common.RenderJSON(w, r, result)
}
