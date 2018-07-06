package CRUD

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/flights/models"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/common"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"net/http"
	"strings"
	"fmt"
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
	request := "SELECT * FROM flights WHERE "
	count := 0
	validKeys := []string{"id", "departure_city", "departure_time", "departure_date", "arrival_city", "arrival_time", "arrival_date", "price"}
	for key, value := range params {
		for _, keys := range validKeys {
			if key == keys {
				count++
			}
		}
		if count == 0 {
			common.SendError(w, r, 400, "ERROR: Invalid request", nil)
			fmt.Println("error")
			return
		}

		switch key {
		case "departure_city", "arrival_city":
			if len(value) > 1 {
				request += key + " IN ("
				for i, v := range value {
					v = "'" + v + "'"
					request += v
					if i < len(value)-1 {
						request += ", "
					}
				}
				request += ") AND "
			} else {
				value[0] = "'" + value[0] + "'"
				request += key + "=" + value[0] + " AND "
			}
		case "departure_time", "departure_date", "arrival_time", "arrival_date", "price":
			if len(value) > 1 {
				request += key + " BETWEEN " + value[0] + " AND " + value[1] + " AND "
			} else {
				request += key + "=" + value[0] + " AND "
			}
		default:
			request += key + "=" + value[0] + " AND "
		}
		count = 0
	}

	words := strings.Fields(request)

	if words[len(words)-1] == "AND" || words[len(words)-1] == "WHERE" {
		words[len(words)-1] = ""
	}

	request = strings.Join(words, " ")
	request += ";"

	museums, err := models.GetByRequest(request)
	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't find flights with such parameters", err)
		return
	}
	common.RenderJSON(w, r, museums)
}
