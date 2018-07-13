package hotels

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

	hotelID, err := uuid.FromString(r.Form.Get("hotel_id"))
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong hotelID (can't convert string to uuid)", err)
		return
	}

	tripID, err := uuid.FromString(r.Form.Get("trip_id"))
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong tripID (can't convert string to uuid)", err)
		return
	}

	err = models.AddHotelToTrip(tripID, hotelID)
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Can't add new hotel to trip", err)
		return
	}
	common.RenderJSON(w, r, success{Status: "201 Created"})
}

func GetByTripHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	tripID, err := uuid.FromString(params["id"])
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong tripID (can't convert string to uuid)", err)
		return
	}

	hotels, err := models.GetHotelsByTrip(tripID)
	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't get hotels by tripID", err)
		return
	}
	common.RenderJSON(w, r, hotels)
}

func GetByRequestHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	hotels, err := models.GetHotelsByRequest(params)
	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't find hotels with such parameters", err)
		return
	}
	common.RenderJSON(w, r, hotels)
}
