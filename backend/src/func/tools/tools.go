package tools

import (
	"net/http"
	"taas/utils"
	"encoding/json"
	"io/ioutil"
	"fmt"
)

func GetTeacherCourseInfo(w http.ResponseWriter, r *http.Request){
	ret := make(map[string]interface{})	
	type Info struct {
		Tid       string `json:"tid"`
		Term      string `json:"term"`
	}
	arr, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		utils.Response(&ret, &w, err.Error())
		return
	}
	var info Info
	err = json.Unmarshal(arr, &info)
	if err != nil {
		utils.Response(&ret, &w, err.Error())
		return
	}
	// --- get json
	fmt.Println(info.Tid, info.Term)
	c,err := utils.QueryTeacherCourseInfo(info.Term, info.Tid)
	if err != nil {
		utils.Response(&ret, &w, err.Error())
		return
	}
	for key, val := range c {
		ret[key] = val
	}
	utils.Response(&ret, &w, "ok")
}

func GetTidTname(w http.ResponseWriter, r *http.Request) {
	ret := make(map[string]interface{})
	c, err := utils.QueryTidTname()
	if err != nil {
		utils.Response(&ret, &w, err.Error())
		return
	}
	for key, val := range c {
		ret[key] = val
	}
	utils.Response(&ret, &w, "ok")
}

func GetDepartment(w http.ResponseWriter, r *http.Request) {
	ret := make(map[string]interface{})
	c, err := utils.GetDepartment()
	if err != nil {
		utils.Response(&ret, &w, err.Error())
		return
	}
	for key, val := range c {
		ret[key] = val
	}
	utils.Response(&ret, &w, "ok")
}

func GetCurrentTerm(w http.ResponseWriter, r *http.Request) {
	term, err := utils.GetCurrentTerm()
	ret := make(map[string]interface{})
	ret["term"] = term
	if err != nil {
		utils.Response(&ret, &w, err.Error())
		return
	}
	utils.Response(&ret, &w, "ok")
}
