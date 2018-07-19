package restaurants

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/models"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/common"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"net/http"
)

type successAdd struct {
	Status string `json:"message"`
}

type successDelete struct {
	Status string `json:"message"`
}

func Get(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	restaurants, err := models.GetRestByQuery(params)
	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't find any restaurants", err)
		return
	}

	common.RenderJSON(w, r, restaurants)
}

func SaveRest(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Can't parse POST Body", err)
		return
	}

	restaurantID, err := uuid.FromString(r.Form.Get("restaurant_id"))
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong restaurant ID (can't convert string to uuid)", err)
		return
	}

	tripID, err := uuid.FromString(r.Form.Get("trip_id"))
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong trip ID (can't convert string to uuid)", err)
		return
	}

	err = models.SaveRest(tripID, restaurantID)
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Can't add new restaurant to trip", err)
		return
	}

	common.RenderJSON(w, r, successAdd{Status: "201 Created"})
}

func Delete(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	itemID, err := uuid.FromString(params["id"])

	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong item ID (can't convert string to uuid)", err)
		return
	}

	err = models.DeleteRestFromDB(itemID)

	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't delete this item", err)
		return
	}

	common.RenderJSON(w, r, successDelete{Status: "204 No Content"})
}

func GetRestFromTrip(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	tripID, err := uuid.FromString(params["id"])
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong trip ID (can't convert string to uuid)", err)
		return
	}

	trains, err := models.GetRestFromTrip(tripID)
	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't get restaurants by trip ID", err)
		return
	}

	common.RenderJSON(w, r, trains)
}
