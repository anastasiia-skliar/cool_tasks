//Package hotels implements hotel handlers
package hotels

import (
	"encoding/json"
	"github.com/Nastya-Kruglikova/cool_tasks/src/model"
	"github.com/Nastya-Kruglikova/cool_tasks/src/service/common"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"net/http"
)

type success struct {
	Status string `json:"message"`
}

type TripHotel struct {
	HotelID string `json:"hotel_id"`
	TripID  string `json:"trip_id"`
}

//AddHotelToTripHandler is a handler for adding Hotel to Trip
func AddHotelToTripHandler(w http.ResponseWriter, r *http.Request) {
	var newHotel TripHotel

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newHotel)
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Can't decode JSON POST Body", err)
		return
	}

	hotelID, err := uuid.FromString(newHotel.HotelID)
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong hotelID (can't convert string to uuid)", err)
		return
	}

	tripID, err := uuid.FromString(newHotel.TripID)
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong tripID (can't convert string to uuid)", err)
		return
	}

	err = model.AddToTrip(tripID, hotelID,model.Hotel{})
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

	hotels, err := model.GetFromTrip(tripID,model.Hotel{})
	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't get hotels by tripID", err)
		return
	}
	common.RenderJSON(w, r, hotels)
}

//GetHotelsHandler is a handler for getting Hotels from Trip by request
func GetHotelsHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	hotels, err := model.GetFromTripWithParams(params,model.Hotel{})
	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't find hotels with such parameters", err)
		return
	}
	common.RenderJSON(w, r, hotels)
}
