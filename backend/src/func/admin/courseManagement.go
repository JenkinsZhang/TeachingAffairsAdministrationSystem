package admin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"taas/utils"
)

func CourseManagement(w http.ResponseWriter, r *http.Request) {
	type Info struct {
		Op     string `json:"op"`
		Did    string `json:"did"`
		Tid    string `json:"tid"`
		Cid    string `json:"cid"`
		Credit string `json:"credit"`
		Term   string `json:"term"`
		Cname  string `json:"cname"`
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
	if len(r.Header["Authorization"]) == 0 {
		w.WriteHeader(http.StatusOK)
		return
	}

	ret := make(map[string]interface{})
	ret["message"] = "ok"
	token := r.Header["Authorization"][0]
	_, err = utils.CheckToken(token)
	if err != nil {
		ret["message"] = "invalid token"
		utils.Response(ret, w)
		return
	}
	// ----
	if r.Method == "GET" {
		ret["term"] = utils.QueryTerm()
	} else if r.Method == "POST" {
		if info.Op == "select" {
			c := utils.QueryCourseWithTerm(info.Term)
			for key, val := range c {
				ret[key] = val
			}
		} else if info.Op == "add" {
			utils.InsertCourse(utils.Struct2Map(info))
		} else if info.Op == "modify" {

		} else if info.Op == "delete" {

		} else {
			ret["message"] = "fail"
		}
	}
	utils.Response(ret, w)
}
