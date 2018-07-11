package hotels
import (
	sq "github.com/Masterminds/squirrel"
	"github.com/Nastya-Kruglikova/cool_tasks/src/models"
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/common"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"net/http"
)

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
		case "id", "departure_time", "departure_date", "arrival_time", "arrival_date", "price":
			if len(params[k]) == 2 {
				cond = append(cond, sq.And{sq.GtOrEq{k: v[0]}, sq.LtOrEq{k: v[1]}})
			} else {
				cond = append(cond, sq.Eq{k: v[0]})
			}
		case "departure_city", "arrival_city":
			cond = append(cond, sq.Eq{k: v[0]})
		case "MAX_PRICE":
			cond = append(cond,sq.Eq{k:v[0]})
		default:
			common.SendError(w, r, 400, "ERROR: Empty or invalid req", nil)
		}
	}

	request, _, _ = hotels.Where(cond).ToSql()

	if len(params) == 0 {
		request = "SELECT * FROM hotels;"
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

	trainID, err := uuid.FromString(r.Form.Get("train_id"))
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong train ID (can't convert string to uuid)", err)
		return
	}

	tripID, err := uuid.FromString(r.Form.Get("trip_id"))
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong trip ID (can't convert string to uuid)", err)
		return
	}

	err = models.AddHotelToTrip(trainID, tripID)
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Can't add new train to trip", err)
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

	trains, err := models.GetHotelFromTrip(tripID)
	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't get trains by trip ID", err)
		return
	}

	common.RenderJSON(w, r, trains)
}

