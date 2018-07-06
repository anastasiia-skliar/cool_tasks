package museums

import (
	"net/http"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/common"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"strings"
)

type successCreate struct {
	Status string `json:"status"`
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

func GetMuseumsByRequestHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	request := "SELECT * FROM museums WHERE "
	count := 0
	validKeys := []string{"id", "name", "location", "price", "museum_type", "opened_at", "closed_at"}
	for key, value := range params {
		for _, keys := range validKeys {
			if key == keys {
				count++
			}
		}
		if count == 0 {
			//common.SendError(w, r, 400, "ERROR: Invalid request", nil)
			return
		}
		switch key {
		case "name", "location", "museum_type":
			value[0] = "'" + value[0] + "'"
			request += key + "=" + value[0] + " AND "
		case "price", "opened_at", "closed_at":
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
	museums, err := GetMuseumsByRequest(request)
	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't find museums with such parameters", err)
		return
	}
	common.RenderJSON(w, r, museums)
}
