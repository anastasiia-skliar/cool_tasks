package models

import (
	"database/sql"
	"fmt"
	. "github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/satori/go.uuid"
	"net/url"
	sq "gopkg.in/Masterminds/squirrel.v1"
)

const (
	datalocation   = "restaurants"
	getter         = "SELECT * FROM %s"
	create         = "INSERT INTO %s (%s) VALUES (%s) RETURNING id"
	getByParameter = "WHERE %s = $1"
	addParam       = " AND %s = $%d"
	addOr = " OR %s = $%d"
	deleteTempl    = "DELETE FROM %s WHERE id = $1"
)

var deleteRequest string

//Restaurant representation in DB
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

func recGen(params map[string][]string) string {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	query, res, err := psql.Select("*").From(datalocation).Where(Lt).ToSql()
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println(res)
	/*base := fmt.Sprintf(getter, datalocation)
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
	*/
	fmt.Println(query)
	return query
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
var AddRestaurant = func(item Restaurant) (Restaurant, error) {
	err := DB.QueryRow(create, item.Name, item.Location, item.Stars, item.Prices, item.Description).Scan(&item.ID)
	return item, err
}

//GetTask used for getting task from DB
var GetRestaurantsByID = func(id uuid.UUID) (Restaurant, error) {
	var item Restaurant
	params:=make(map[string]string)
	key :=make([]string,0)
	key=append(key, id.String())
params["id"]=key
	err := DB.QueryRow(recGen(params)).Scan(&item.ID, &item.Name, &item.Location, &item.Stars, &item.Prices, &item.Description)
	return item, err
}

//DeleteTask used for deleting task from DB
var DeleteRestaurantsFromDB = func(id uuid.UUID) error {
	_, err := DB.Exec(deleteRequest, id)
	return err
}

//GetTasks used for getting tasks from DB

var GetRestaurantsByQuery = func(query url.Values) ([]Restaurant, error) {
	rows, err := DB.Query(recGen(query))

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
