//Package hotels implements hotel handlers
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

//AddHotelToTripHandler is a handler for adding Hotel to Trip
func AddHotelToTripHandler(w http.ResponseWriter, r *http.Request) {
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

	err = models.AddToTrip(hotelID, tripID, models.Hotel{})
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Can't add new hotel to trip", err)
		return
	}
	common.RenderJSON(w, r, success{Status: "201 Created"})
}

//GetHotelsByTripHandler is a handler for getting Hotels from Trip
func GetHotelsByTripHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	tripID, err := uuid.FromString(params["id"])
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong tripID (can't convert string to uuid)", err)
		return
	}

	hotels, err := models.GetFromTrip(tripID, models.Hotel{})
	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't get hotels by tripID", err)
		return
	}
	common.RenderJSON(w, r, hotels)
}

//GetHotelsHandler is a handler for getting Hotels from Trip by request
func GetHotelsHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	hotels, err := models.GetData(params, models.Hotel{})
	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't find hotels with such parameters", err)
		return
	}
	common.RenderJSON(w, r, hotels)
}
