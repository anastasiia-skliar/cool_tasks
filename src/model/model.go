package model

import (
	"fmt"
	"github.com/Nastya-Kruglikova/cool_tasks/src/database"
	"github.com/satori/go.uuid"
	"log"
	"net/url"
	"reflect"
	"strings"
	"time"
)

//Event is a representation of Event table in DB
type Event struct {
	ID       uuid.UUID
	Title    string
	Category string
	Town     string
	Date     time.Time
	Price    int
}

//Flight is a representation of Flight table in DB
type Flight struct {
	ID             uuid.UUID
	Departure_city string
	Departure      time.Time
	Arrival_city   string
	Arrival        time.Time
	Price          int
}

//Museum is a representation of Museum table in DB
type Museum struct {
	ID              uuid.UUID
	Name            string
	Location        string
	Price           int
	Opened_at       time.Time
	Closed_at       time.Time
	Museum_type     string
	Additional_info string
}

//Train representation in DB
type Train struct {
	ID             uuid.UUID
	Departure      time.Time
	Arrival        time.Time
	Departure_city string
	Arrival_city   string
	Train_type     string
	Car_type       string
	Price          int
}

//Hotel representation in DB
type Hotel struct {
	ID         uuid.UUID
	Name       string
	Class      string
	Capacity   int
	Rooms_left int
	Floors     int
	Price      int
	City_name  string
	Address    string
}

//Restaurant representation in DB
type Restaurant struct {
	ID          uuid.UUID
	Name        string
	Location    string
	Stars       int
	Prices      int
	Description string
}

var AddToTrip = func(dataID uuid.UUID, tripID uuid.UUID, dataSource interface{}) error {
	_, err := database.DB.Exec(GenerateQueryAdd(dataSource), dataID, tripID)
	return err
}

var GetFromTrip = func(tripID uuid.UUID, obj interface{}) (interface{}, error) {
	rows, err := database.DB.Query(generateQueryGet(obj), tripID)
	if err != nil {
		return nil, err
	}

	cols, err := rows.Columns()
	rowType := reflect.ValueOf(obj).Type()
	slicePtrVal := reflect.New(reflect.SliceOf(rowType))
	sliceVal := reflect.Indirect(slicePtrVal)

	for rows.Next() {
		var row = make([]interface{}, len(cols))
		var rowp = make([]interface{}, len(cols))
		for i := 0; i < len(cols); i++ {
			rowp[i] = &row[i]
		}
		val := reflect.ValueOf(obj)
		vp := reflect.New(val.Type())

		rows.Scan(rowp...)

		var v interface{}

		for i, col := range cols {
			fieldName := strings.ToUpper(col[0:1]) + strings.ToLower(col[1:])
			if fieldName == "Id" {
				fieldName = strings.ToUpper(fieldName)
			}
			v = row[i]
			structField := vp.Elem().FieldByName(fieldName)

			condition := structField.Type().Name()
			if condition == "UUID" {
				s := string(reflect.ValueOf(row[i]).Bytes()[:])
				v, err = uuid.FromString(s)
				if err != nil {
					log.Println(err)
				}
			} else if condition == "int" {
				v = int(reflect.ValueOf(v).Int())
			}
			vp.Elem().FieldByName(fieldName).Set(reflect.ValueOf(v))
		}

		sliceVal.Set(reflect.Append(sliceVal, vp.Elem()))
	}
	return sliceVal.Interface(), nil
}

var GetFromTripWithParams = func(params url.Values, obj interface{}) (interface{}, error) {
	objType := reflect.TypeOf(obj)
	name := strings.ToLower(objType.Name())
	pluralName := name + "s"

	var stringArgs []string
	var numberArgs []string

	switch obj.(type) {

	case Event:
		stringArgs = []string{"title", "category", "town"}
		numberArgs = []string{"price", "date"}

	case Flight:
		stringArgs = []string{"departure_city", "arrival_city"}
		numberArgs = []string{"price", "departure_time", "arrival_time", "departure_date", "arrival_date"}
	case Museum:
		stringArgs = []string{"name", "location", "museum_type"}
		numberArgs = []string{"price", "opened_at", "closed_at"}

	case Train:
		stringArgs = []string{"departure_city", "arrival_city"}
		numberArgs = []string{"price", "departure_time", "arrival_time", "departure_date", "arrival_date"}

	case Restaurant:
		stringArgs = []string{"id", "name", "location"}
		numberArgs = []string{"stars", "prices"}

	case Hotel:
		stringArgs = []string{"name", "city_name", "address"}
		numberArgs = []string{"class", "capacity", "rooms_left", "floors", "price"}
	}

	request, args, err := SQLGenerator(pluralName, stringArgs, numberArgs, params)

	if err != nil {
		return nil, err
	}

	rows, err := database.DB.Query(request, args...)
	if err != nil {
		return nil, err
	}

	cols, err := rows.Columns()
	rowType := reflect.ValueOf(obj).Type()
	slicePtrVal := reflect.New(reflect.SliceOf(rowType))
	sliceVal := reflect.Indirect(slicePtrVal)

	for rows.Next() {
		var row = make([]interface{}, len(cols))
		var rowp = make([]interface{}, len(cols))
		for i := 0; i < len(cols); i++ {
			rowp[i] = &row[i]
		}
		val := reflect.ValueOf(obj)
		vp := reflect.New(val.Type())

		rows.Scan(rowp...)

		var v interface{}

		for i, col := range cols {
			fieldName := strings.ToUpper(col[0:1]) + strings.ToLower(col[1:])
			if fieldName == "Id" {
				fieldName = strings.ToUpper(fieldName)
			}
			v = row[i]
			structField := vp.Elem().FieldByName(fieldName)

			condition := structField.Type().Name()
			if condition == "UUID" {
				s := string(reflect.ValueOf(row[i]).Bytes()[:])
				v, err = uuid.FromString(s)
				if err != nil {
					fmt.Println(err)
				}
			} else if condition == "int" {
				v = int(reflect.ValueOf(v).Int())
			}
			vp.Elem().FieldByName(fieldName).Set(reflect.ValueOf(v))

		}
		sliceVal.Set(reflect.Append(sliceVal, vp.Elem()))
	}
	return sliceVal.Interface(), nil
}

func GenerateQueryAdd(dataSource interface{}) string {
	dataType := reflect.TypeOf(dataSource)
	var query = "INSERT INTO trips_" + strings.ToLower(dataType.Name()) + "s" + " (" + strings.ToLower(dataType.Name()) + "_id, trip_id) VALUES ($1, $2)"
	return query
}

func generateQueryGet(dataSource interface{}) string {
	dataType := reflect.TypeOf(dataSource)
	name := strings.ToLower(dataType.Name())
	pluralName := name + "s"
	var query = "SELECT " + pluralName + ".* FROM " + pluralName + " INNER JOIN trips_" + pluralName + " ON " + pluralName + ".id=trips_" + pluralName + "." + name + "_id AND trips_" + pluralName + ".trip_id=$1"
	return query
}
