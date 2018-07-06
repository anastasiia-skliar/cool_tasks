package restaurants

import (
	"github.com/Nastya-Kruglikova/cool_tasks/src/services/common"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	"net/http"
	"strconv"
)

type successCreate struct {
	Status string      `json:"message"`
	Result Restaurant `json:"result"`
}

type successDelete struct {
	Status string      `json:"message"`
}

func Get(w http.ResponseWriter, r *http.Request) {
	query:=r.URL.Query()
	if val, ok := query["id"]; ok {
		id, err:=uuid.FromString(val[0])
		if err != nil {
			common.SendNotFound(w, r, "ERROR: Invalid ID", err)
			return
		}
		items, err := getByID(id)

		if err != nil {
			common.SendNotFound(w, r, "ERROR: Can't get items", err)
			return
		}

		common.RenderJSON(w, r, items)
	}

	paramNames:=make([]string, 0)
	paramVals:=make([]string,0)
	var usedVal string
	for key, value := range query {
	paramNames=	 append(paramNames, key)
		if len(value)>0 {
			usedVal=value[0]
		}else {
			continue
		}
	paramVals=	append(paramVals,usedVal)
	}
	//MAGIC BEGINS!!!
	s := make([]interface{}, len(paramVals))
	for i, v := range paramVals {
		s[i] = v
	}
	items, err := getByParams(paramNames, s...)
//MAGIC ENDS!!!
	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't get items", err)
		return
	}

	common.RenderJSON(w, r, items)
}

func PostRestaurant(w http.ResponseWriter, r *http.Request) {

	var newItem Restaurant
	var resultItem Restaurant

	err := r.ParseForm()

	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Can't parse POST Body", err)
		return
	}



	newItem.Name = r.Form.Get("name")
	newItem.Location=r.Form.Get("location")
	newItem.Description = r.Form.Get("description")
	newItem.Prices, err=	strconv.Atoi(r.Form.Get("prices"))
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Invalid prices field", err)
		return
	}
	newItem.Stars,err=	strconv.Atoi(r.Form.Get("stars"))
	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Invalid stars field", err)
		return
	}

	resultItem, err = addRestaurant(newItem)

	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Can't add new item", err)
		return
	}

	common.RenderJSON(w, r, successCreate{Status: "201 Created", Result: resultItem})
}

func Delete(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	itemID, err := uuid.FromString(params["id"])

	if err != nil {
		common.SendBadRequest(w, r, "ERROR: Wrong item ID (can't convert string to uuid)", err)
		return
	}

	err = deleteFromDB(itemID)

	if err != nil {
		common.SendNotFound(w, r, "ERROR: Can't delete this item", err)
		return
	}

	common.RenderJSON(w, r, successDelete{Status: "204 No Content"})
}
