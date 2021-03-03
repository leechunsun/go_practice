package main

import "fmt"

type MyStruct struct {
	Name string `ca:"name"`
	Age int	`ca:"age"`
}

func main() {
	//a := make(map[string]string)
	//a["name"] = "chunsheng"
	//a["age"] = "17"
	//m := &MyStruct{}
	//reflectVal := reflect.ValueOf(m)
	//reflectTyp := reflect.ValueOf(m).Elem().Type()
	//for i:=0; i<reflectTyp.NumField();i++{
	//	fliedTagName := reflectTyp.Field(i).Tag.Get("ca")
	//	if val, ok:= a[fliedTagName];ok{
	//		fmt.Println(reflectVal.Elem().Field(i).Type(), " : ", val)
	//		var v reflect.Value
	//		switch reflectVal.Elem().Field(i).Type().Kind() {
	//			case reflect.String:
	//				v = reflect.ValueOf(interface{}(val).(string))
	//			case reflect.Int:
	//				ix , err := strconv.Atoi(val)
	//				if err != nil{
	//					ix = 0
	//				}
	//				fmt.Println(ix)
	//				v = reflect.ValueOf(ix)
	//		}
	//		reflectVal.Elem().Field(i).Set(v)
	//		// reflectVal.Elem().Field(i).Set(reflect.ValueOf(val))
	//	}
	//}
	//fmt.Printf("%#v", m)
	//fmt.Println(m.Age)

	//fmt.Println(strconv.FormatFloat(float64(12345.5887), 'g', -1, 64))

	a := make(map[string]string, 0)
	a["aaa"] = "zzz"
	a["a"] = "z"
	a["aa"] = "zz"
	a["aaaa"] = "zzzz"
	for key, val := range a{
		if val == "zzz"{
			delete(a, key)
		}
	}
	b := []int{1,2,3,4,5,6,6,7,7,8,8,9}
	for i, val := range b{
		fmt.Println(i, val)
		if val == 8 {
			b = append(b[:i], b[i+1:]...)
		}
	}
	fmt.Println(b)
}





