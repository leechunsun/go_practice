package main

import (
	"fmt"
	"reflect"
)


type MyStructs struct {
	Name string `ca:"name"`
	Age int	`ca:"age"`
}

func (m *MyStructs) GGG () string {
	return "zzzz"
}

func main() {
	g := MyStructs{}
	mp := make(map[string]interface{}, 5)
	mp["1"] = "1111"
	mp["2"] = 2222
	mp["3"] = 3333
	res := reflect.ValueOf(g).MethodByName("GGG").Call([]reflect.Value{})
	for _, item := range res{
		fmt.Println(item.Interface())
	}


}
