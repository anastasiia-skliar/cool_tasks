//Package restaurants implements restaurant handlers
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

//GetRestaurantHandler used for getting restaurants
func GetRestaurantHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	if val, ok := query["id"]; ok {
		_, err := uuid.FromString(val[0])
		if err != nil {
			common.SendNotFound(w, r, "ERROR: Invalid ID", err)
			return
		}
		restaurants, err := models.GetRestaurants(query)

		if err != nil {
			common.SendNotFound(w, r, "ERROR: Can't get restaurant", err)
			return
		}
		restaurant := restaurants[0]
		common.RenderJSON(w, r, restaurant)
	}

	restaurants, err := models.GetRestaurants(query)

	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't get restaurants", err)
		return
	}

	common.RenderJSON(w, r, restaurants)
}

//AddRestaurantToTrip saves Restaurant to Trip
func AddRestaurantToTripHandler(w http.ResponseWriter, r *http.Request) {
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

	err = models.AddRestaurantToTrip(tripID, restaurantID)
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Can't add new restaurant to trip", err)
		return
	}

	common.RenderJSON(w, r, successAdd{Status: "201 Created"})
}

//DeleteRestaurantHandler deletes Restaurant from DB
func DeleteRestaurantHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	itemID, err := uuid.FromString(params["id"])

	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong item ID (can't convert string to uuid)", err)
		return
	}

	err = models.DeleteRestaurant(itemID)

	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't delete this item", err)
		return
	}

	common.RenderJSON(w, r, successDelete{Status: "204 No Content"})
}

//GetRestaurantFromTrip gets Restaurant from Trip by tripID
func GetRestaurantFromTrip(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	tripID, err := uuid.FromString(params["id"])
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong trip ID (can't convert string to uuid)", err)
		return
	}

	items, err := models.GetRestaurantsFromTrip(tripID)
	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't get restaurants by trip ID", err)
		return
	}

	common.RenderJSON(w, r, items)
}
