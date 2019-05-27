package student

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"taas/utils"
)

func CourseCalendar(w http.ResponseWriter, r *http.Request) {
	ret := make(map[string]interface{})

	// --- token 检查
	claims, err := utils.PreCheck(r)
	if err != nil {
		utils.Response(&ret, &w, err.Error())
		return
	}
	id := claims["id"].(string)

	// ---
	if r.Method == "GET" {
		ret["term"], err = utils.GetAllTerms()
		if err != nil {
			utils.Response(&ret, &w, err.Error())
			return
		}
	} else if r.Method == "POST" {
		type Info struct {
			Term string `json:"term"`
			Cid  string `json:"cid"`
			Op   string `json:"op"`
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
		if info.Term == "" {
			info.Term, _ = utils.GetCurrentTerm()
		}
		// --- get json

		c := make(map[string][]string)
		tmp := make(map[string][]string)
		// 先删后查
		if info.Op == "delete" {
			term, err := utils.GetCurrentTerm()
			if err != nil {
				utils.Response(&ret, &w, "please set current term first.")
				return
			}
			if term != info.Term {
				utils.Response(&ret, &w, "you can't modify the courses of past terms.")
				return
			}
			osc, err := utils.IfOpenSelectCourse()
			if err != nil {
				utils.Response(&ret, &w, err.Error())
				return
			}
			if osc == "close" {
				utils.Response(&ret, &w, "Lesson selection time is not yet available.")
				return
			}
			err = utils.DeleteCourseCalendar(id, info.Cid, info.Term)
			if err != nil {
				utils.Response(&ret, &w, err.Error())
				return
			}
		}
		tmp, err = utils.QueryStuCourses(id, info.Term)
		if err != nil {
			utils.Response(&ret, &w, err.Error())
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
	}
	utils.Response(&ret, &w, "ok")
}
