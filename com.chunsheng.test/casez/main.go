package main

import (
	"fmt"
	"reflect"
	"time"
)

type BaseStruct struct {
	Name string `lcs:"name"`
	Age int64 `lcs:"age"`
	Addr string `lcs:"addr"`
	Birth time.Time `lcs:"birth"`

}


func main() {
	obj := BaseStruct{}
	coreMap := make(map[string]interface{}, 10)
	coreMap["name"] = "1234"
	coreMap["age"] = int64(10)
	coreMap["birth"] = time.Now()
	objval := reflect.ValueOf(&obj).Elem()
	for i:=0;i<objval.NumField();i++{
		tag := objval.Type().Field(i).Tag.Get("lcs")
		// fieldType := objval.Field(i).Type().Name()
		fieldTyp := objval.Type().Field(i).Type.String()
		val := coreMap[tag]
		switch fmt.Sprintf("%T", val){
		case fieldTyp:
			objval.Field(i).Set(reflect.ValueOf(val))
		default:
			fmt.Println("zzzzz")
		}
	}
	fmt.Println(obj)

}
