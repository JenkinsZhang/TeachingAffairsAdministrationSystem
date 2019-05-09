package student

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"utils"
)

func CourseQuery(w http.ResponseWriter, r *http.Request) {
	type Info struct {
		Cid       string `json:"cid"`
		Tid       string `json:"tid"`
		Cname     string `json:"cname"`
		Tname     string `json:"tname"`
		Term      string `json:"term"`
		ClassTime string `json:"ClassTime"`
		Op        string `json:"op"`
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

	ret := make(map[string]interface{})
	ret["message"] = "ok"
	token := r.Header["Authorization"][0]
	claims, err := utils.CheckToken(token)
	id := claims["id"].(string)
	if err != nil {
		ret["message"] = "invalid token"
		utils.Response(ret, w)
		return
	}
	// ----
	c := make(map[string][]string)

	if info.Op == "query" {
		if info.Cid != "" {
			c = utils.QueryCourseWithCid(info.Cid)
		} else if info.Cname != "" {
			c = utils.QueryCourseWithCname(info.Cname)
		} else if info.Tid != "" {
			c = utils.QueryCourseWithTid(info.Tid)
		} else if info.Tname != "" {
			c = utils.QueryCourseWithTname(info.Tname)
		} else {
			c = utils.QueryAllCourses()
		}
		for key, val := range c {
			ret[key] = val
		}
	} else if info.Op == "select" {
		tmp := utils.Struct2Map(info)
		tmp["Id"] = id
		utils.InsertCourse(tmp)
	}

	utils.Response(ret, w)
}
