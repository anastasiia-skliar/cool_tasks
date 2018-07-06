package handlers

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/common"
	"github.com/Nastya-Kruglikova/cool_tasks/src/trains/models"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"net/http"
)

const (
	AND         = " AND "
	BETWEEN     = " BETWEEN "
	equals      = " = "
	selectWhere = "SELECT * FROM trains WHERE "
	selectAll   = "SELECT * FROM trains;"
	qm          = "'"
)

type successAdd struct {
	Status string       `json:"message"`
	Result models.Train `json:"result"`
}

func GetTrains(w http.ResponseWriter, r *http.Request) {
	var request string
	req := selectWhere
	params := r.URL.Query()

	for k := range params {
		switch k {
		case "id", "departure_time", "departure_date", "arrival_time", "arrival_date", "price":
			if len(params[k]) == 2 {
				req += k + BETWEEN + qm + params[k][0] + qm + AND + qm + params[k][1] + qm + AND
			} else {
				req += k + equals + qm + params[k][0] + qm + AND
			}
		case "departure_city", "arrival_city":
			req += k + equals + qm + params[k][0] + qm + AND
		default:
			common.SendError(w, r, 400, "ERROR: Empty or invalid req", nil)
		}
	}
	request = req[:len(req)-5] + ";"

	if len(params) == 0 {
		request = selectAll
	}

	trains, err := models.GetTrains(request)
	if err != nil {
		common.SendError(w, r, 400, "ERROR: Empty or invalid req", nil)
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
		common.SendBadRequest(w, r, "ERROR: Wrong flight ID (can't convert string to uuid)", err)
		return
	}

	tripID, err := uuid.FromString(r.Form.Get("trip_id"))
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong trip ID (can't convert string to uuid)", err)
		return
	}

	err = models.SaveTrain(trainID, tripID)
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Can't add new flight to trip", err)
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
