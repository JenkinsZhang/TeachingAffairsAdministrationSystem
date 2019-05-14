package utils

import (
	"encoding/json"
	"net/http"
	"reflect"
)

func Struct2Map(obj interface{}) map[string]string {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]string)
	for i := 0; i < t.NumField(); i++ {
		// fmt.Println(t.Field(i).Name, v.Field(i).Interface().(string))
		data[t.Field(i).Name] = v.Field(i).Interface().(string)
	}
	return data
}

func Response(ret *map[string]interface{}, w *http.ResponseWriter, msg string) {
	(*ret)["message"] = msg
	jtmp, err := json.Marshal(ret)
	if err != nil {
		(*ret)["message"] = err.Error()
	}
	(*w).Write(jtmp)
	(*w).WriteHeader(http.StatusOK)
}
