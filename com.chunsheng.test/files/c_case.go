package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

type My struct {
	Name string `config:"name"`
	Age int `config:"age"`
	Addr string `config:"addr"`
}

func (m *My) callTest(){
	fmt.Println("you are success")
}

func ReadFile(configPath string) (conMap map[string]string) {
	content, e := ioutil.ReadFile(configPath)
	if e != nil{
		panic(e)
	}
	conMap = make(map[string]string)
	conA := strings.Split(string(content), "\n")
	for _, st := range conA{
		st_tri := strings.TrimSpace(st)
		conList := strings.Split(st_tri, "=")
		if len(conList) < 2{
			continue
		}
		conMap[conList[0]] = conList[1]
	}
	return
}


func ConfigInsert(configPath string, ca interface{}){
	conMap := ReadFile(configPath)
	fmt.Println(conMap)
	rfe := reflect.TypeOf(ca).Elem()
	rfv := reflect.ValueOf(ca).Elem()
	for i:=0;i < rfe.NumField(); i++ {
		key := rfe.Field(i).Tag.Get("config")
		if val, ok := conMap[key];ok{
			switch rfv.Field(i).Kind() {
			case reflect.String:
				rfv.Field(i).SetString(val)

			case reflect.Int:
				v, _:= strconv.ParseInt(val, 10, 64)
				rfv.Field(i).SetInt(v)
			}
		}
	}
}



func main() {
	file_path := "/Users/zrb/Documents/codes/go/src/com.chunsheng.test/files/case.conf"
	my := &My{}
	ConfigInsert(file_path, my)
	fmt.Printf("%#v", my)
}
