package main

import (
	"fmt"
	"reflect"
)

type Myc struct {
	Name string `config:"name"`
	Age int `config:"age"`
	Addr string `config:"addr"`
}

func (m *Myc) CallTest(k string){
	fmt.Printf("you are success: %s", k)
}


func main() {
	my := &Myc{}
	var x reflect.Value
	rfv := reflect.ValueOf(my)
	rft := reflect.TypeOf(my)
	fmt.Println(rfv.NumMethod())
	for i := 0;i< rfv.NumMethod();i++{
		x = rfv.Method(i)
		x.Call([]reflect.Value{reflect.ValueOf("xx")})
	}

	for j:=0;j< rft.NumMethod();j++{
		z := rft.Method(j).Func
		z.Call([]reflect.Value{reflect.ValueOf("xx")})
	}
}
