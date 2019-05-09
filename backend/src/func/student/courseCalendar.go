package student

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"utils"
)

func CourseCalendar(w http.ResponseWriter, r *http.Request) {
	type Info struct {
		Term string `json:"term"`
		Cid  string `json:"cid"`
		Op   string `json:"op"`
	}
	r.ParseForm()
	arr, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "readall: %v\n", err)
		os.Exit(-1)
	}
	var info Info
	err = json.Unmarshal(arr, &info)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unmarshal: %v\n", err)
		os.Exit(-1)
	}
	if r.Method != "POST" || len(r.Header["Authorization"]) == 0 {
		w.WriteHeader(http.StatusOK)
		return
	}
	// ----

	c := make(map[string][]string)
	ret := make(map[string]interface{})
	tmp := make(map[string][]string)

	ret["message"] = "ok"
	token := r.Header["Authorization"][0]
	claims, err := utils.CheckToken(token)
	id := claims["id"].(string)
	if err != nil {
		ret["message"] = "invalid token"
		utils.Response(ret, w)
		return
	}
	// 先删后查
	if info.Op == "delete" {
		utils.DeleteCourse(id, info.Cid, info.Term)
	}
	tmp = utils.QueryStuCourses(id, info.Term)
	if tmp == nil {
		ret["message"] = "invalid token"
		utils.Response(ret, w)
		return
	}
	for key, val := range tmp {
		for _, v := range val {
			c[key] = append(c[key], v)
		}
	}
	for key, val := range c {
		ret[key] = val
	}
	utils.Response(ret, w)

}
