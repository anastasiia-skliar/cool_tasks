package hotels

import (
	sq "github.com/Masterminds/squirrel"
	"github.com/Nastya-Kruglikova/cool_tasks/src/models"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/common"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"log"
	"net/http"
)

const selectAllHotels = "SELECT * FROM hotels;"

type successAdd struct {
	Status string       `json:"message"`
	Result models.Hotel `json:"result"`
}

func GetHotels(w http.ResponseWriter, r *http.Request) {
	var cond sq.And
	var request string
	var err error

	params := r.URL.Query()
	hotels := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).Select("*").From("hotels")

	for k, v := range params {
		switch k {
		case "id", "name", "class", "capacity", "room_left", "floors", "max_price", "city_name", "address":
			if len(params[k]) == 2 {
				cond = append(cond, sq.And{sq.GtOrEq{k: v[0]}, sq.LtOrEq{k: v[1]}})
			} else {
				cond = append(cond, sq.Eq{k: v[0]})
			}
		default:
			common.SendError(w, r, 400, "ERROR: Empty or invalid req", nil)
		}
	}

	request, _, err = hotels.Where(cond).ToSql()
	if err != nil {
		log.Println(err)
	}

	if len(params) == 0 {
		request = selectAllHotels
	}

	result, err := models.GetHotels(request)
	if err != nil {
		common.SendError(w, r, 400, "ERROR: Empty or invalid req", nil)
	}

	common.RenderJSON(w, r, result)
}

func AddHotel(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Can't parse POST Body", err)
		return
	}

	hotelID, err := uuid.FromString(r.Form.Get("hotel_id"))
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong hotel ID (can't convert string to uuid)", err)
		return
	}

	tripID, err := uuid.FromString(r.Form.Get("trip_id"))
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong trip ID (can't convert string to uuid)", err)
		return
	}

	err = models.AddHotelToTrip(hotelID, tripID)
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Can't add new hotel to trip", err)
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

	hotels, err := models.GetHotelFromTrip(tripID)
	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't get hotels by trip ID", err)
		return
	}

	common.RenderJSON(w, r, hotels)
}
