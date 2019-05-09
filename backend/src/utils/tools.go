package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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

func Response(ret map[string]interface{}, w http.ResponseWriter) {
	jtmp, err := json.Marshal(ret)
	if err != nil {
		fmt.Fprintf(os.Stderr, "marshal: %v\n", err)
		os.Exit(-1)
	}
	w.Write(jtmp)
	w.WriteHeader(http.StatusOK)
}
