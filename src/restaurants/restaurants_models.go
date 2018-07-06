package restaurants

import (
	"database/sql"
	"fmt"
	. "github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/satori/go.uuid"
	"net/url"
)

const (
	datalocation   = "restaurants"
	getter         = "SELECT * FROM %s"
	create         = "INSERT INTO restaurants (name, location, stars, prices, description) VALUES ($1, $2, $3, $4, $5) RETURNING id"
	getByParameter = "WHERE %s = $1"
	addParam       = " AND %s = $%d"
	addOr = " OR %s = $%d"
	deleteTempl    = "DELETE FROM %s WHERE id = $1"
)

var deleteRequest string

//Task representation in DB
type Restaurant struct {
	ID          uuid.UUID
	Name        string
	Location    string
	Stars       int
	Prices      int
	Description string
}

func init() {
	deleteRequest = fmt.Sprintf(deleteTempl, datalocation)

}

func recGen(params ...string) string {
	base := fmt.Sprintf(getter, datalocation)
	if len(params) < 1 {
		return base
	}
	paramsCounter := 0
	request := fmt.Sprintf(base+" "+getByParameter, params[paramsCounter])
	paramsCounter++
	for ; paramsCounter < len(params); paramsCounter++ {
		if params[paramsCounter]!=params[paramsCounter-1] {
			request += fmt.Sprintf(addParam, params[paramsCounter], paramsCounter+1)
		}else {
			request += fmt.Sprintf(addOr, params[paramsCounter], paramsCounter+1)
		}

	}
	fmt.Println(request)
	return request
}

func parseResult(rows *sql.Rows) ([]Restaurant, error) {
	res := make([]Restaurant, 0)

	for rows.Next() {
		var item Restaurant
		if err := rows.Scan(&item.ID, &item.Name, &item.Location, &item.Stars, &item.Prices, &item.Description); err != nil {
			return []Restaurant{}, err
		}
		res = append(res, item)
	}
	return res, nil
}

//CreateTask used for creation task in DB
var addRestaurant = func(item Restaurant) (Restaurant, error) {
	err := DB.QueryRow(create, item.Name, item.Location, item.Stars, item.Prices, item.Description).Scan(&item.ID)
	return item, err
}

//GetTask used for getting task from DB
var getByID = func(id uuid.UUID) (Restaurant, error) {
	var item Restaurant
	err := DB.QueryRow(recGen("id"), id).Scan(&item.ID, &item.Name, &item.Location, &item.Stars, &item.Prices, &item.Description)
	return item, err
}

//DeleteTask used for deleting task from DB
var deleteFromDB = func(id uuid.UUID) error {
	_, err := DB.Exec(deleteRequest, id)
	return err
}

//GetTasks used for getting tasks from DB

var getByQuery = func(query url.Values) ([]Restaurant, error) {

	paramNames := make([]string, 0)
	paramVals := make([]string, 0)
	for key, value := range query {
		if len(value) > 0 {
			for _, v:=range value{
				paramNames = append(paramNames, key)
				paramVals = append(paramVals, v)
			}
		} else {
			continue
		}

	}

	s := make([]interface{}, len(paramVals))
	for i, v := range paramVals {
		s[i] = v
	}

	rows, err := DB.Query(recGen(paramNames...), s...)

	if err != nil {
		fmt.Println(err)
		return []Restaurant{}, err
	}
	res, err := parseResult(rows)
	if err != nil {
		fmt.Println(err)
		return []Restaurant{}, err
	}
	return res, nil
}
