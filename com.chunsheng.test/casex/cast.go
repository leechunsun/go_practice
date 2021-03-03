package main

import (
	"fmt"
	"reflect"
)

func CastMapToObject(m map[string]interface{}, obj interface{}, tarBy string){
	rv := reflect.ValueOf(obj)
	if rv.Kind() != reflect.Ptr{
		fmt.Println("类型需要为指针类型。。。")
		return
	}
	rt := reflect.TypeOf(obj).Elem()
	for i := 0; i < rt.NumField(); i++ {
		tarKey := rt.Field(i).Tag.Get(tarBy)
		fieldType := rt.Field(i).Type
		fmt.Println(tarKey, fieldType)
		rv.Elem().Field(i).Set(reflect.ValueOf("eeee"))
	}
}


type My struct {
	Name string
}


func main() {
	my := My{}
	mm := make(map[string]interface{}, 10)
	CastMapToObject(mm, &my, "zzzzz")
	fmt.Println(my)
}
