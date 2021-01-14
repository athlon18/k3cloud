package util

import (
	"encoding/json"
	"fmt"
	_struct "github.com/athlon18/k3cloud/service/struct"
	"reflect"
)

func FNumber(str string) _struct.FNumber {
	fNumber := _struct.FNumber{}
	fNumber.FNumber = str
	return fNumber
}

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func StructJsonMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Tag.Get("json")] = v.Field(i).Interface()
	}
	return data
}

func J(obj interface{}) {
	bb, _ := json.Marshal(obj)
	fmt.Println(string(bb))
}

func ToJson(obj interface{}) string {
	bb, _ := json.Marshal(obj)
	return string(bb)
}
