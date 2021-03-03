package strs

import (
	"fmt"
	"reflect"
)

func String(origin interface{}) string {
	if _, ok := origin.(reflect.Type); ok {
		origin = reflect.ValueOf(origin).Interface()
	}
	switch reflect.TypeOf(origin).Kind() {
	case reflect.String:
		return origin.(string)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%#v", origin)
	case reflect.Ptr:
		return String(reflect.ValueOf(origin).Elem())
	default:
		print("not support type:", reflect.TypeOf(origin).Kind())
		return ""
	}
}
