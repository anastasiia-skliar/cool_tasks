package main

import (
	"reflect"
	"strings"
	"fmt"
	"github.com/Nastya-Kruglikova/cool_tasks/src/models"
)

func generateQueryGet(dataSource interface{})string{
	dataType:=reflect.TypeOf(dataSource)
	name:=strings.ToLower(dataType.Name())
	pluralName:=name+"s"
	var query  ="SELECT "+pluralName+".* FROM "+pluralName+" INNER JOIN trips_"+pluralName+" ON "+pluralName+".id=trips_"+pluralName+"."+name+"_id AND trips_"+pluralName+".trip_id=$1"
	return query
}

func main(){
	fmt.Println(models.GenerateQueryAdd(models.Train{}))
}