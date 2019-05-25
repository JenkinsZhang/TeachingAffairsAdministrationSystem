package admin

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"taas/utils"
)

func CourseManagement(w http.ResponseWriter, r *http.Request) {
	ret := make(map[string]interface{})

	// --- token 检查
	_, err := utils.PreCheck(r)
	if err != nil {
		utils.Response(&ret, &w, err.Error())
		return
	}
	// aid := claims["id"].(string)

	if r.Method == "GET" {
		ret["term"], err = utils.GetAllTerms()
		if err != nil {
			utils.Response(&ret, &w, err.Error())
			return
		}
		ret["course"], err = utils.GetCourse()
		if err != nil {
			utils.Response(&ret, &w, err.Error())
			return
		}
	} else if r.Method == "POST" {
		type Info struct {
			Op        string `json:"op"`
			Did       string `json:"did"`
			Tid       string `json:"tid"`
			Cid       string `json:"cid"`
			Credit    string `json:"credit"`
			Term      string `json:"term"`
			Cname     string `json:"cname"`
			Dname     string `json:"dname"`
			ClassTime string `json:"classTime"`
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
		// getjson ---------------------
		if info.Op == "select" {
			c, err := utils.QueryCourseNumberWithTerm(info.Term)
			if err != nil {
				utils.Response(&ret, &w, err.Error())
				return
			}
			for key, val := range c {
				ret[key] = val
			}
		} else if info.Op == "calendar" {
			c, err := utils.QueryCourseWithTermAndCid(info.Term, info.Cid)
			if err != nil {
				utils.Response(&ret, &w, err.Error())
				return
			}
			for key, val := range c {
				ret[key] = val
			}
		} else if info.Op == "addCourse" {
			err = utils.InsertCourse(utils.Struct2Map(info))
			if err != nil {
				utils.Response(&ret, &w, err.Error())
				return
			}
		} else if info.Op == "deleteCourse" {
			err = utils.DeleteCourse(info.Cid)
			if err != nil {
				utils.Response(&ret, &w, err.Error())
				return
			}
		} else if info.Op == "addCourseSchedule" {
			info.Term, err = utils.GetCurrentTerm()
			if err != nil {
				utils.Response(&ret, &w, err.Error())
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
			err = utils.InsertCourseSchedule(utils.Struct2Map(info))
			if err != nil {
				utils.Response(&ret, &w, err.Error())
				return
			}
		} else if info.Op == "deleteCourseSchedule" {
			info.Term, err = utils.GetCurrentTerm()
			if err != nil {
				utils.Response(&ret, &w, err.Error())
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
			err = utils.DeleteCourseSchedule(info.Cid, info.Tid, info.Term)
			if err != nil {
				utils.Response(&ret, &w, err.Error())
				return
			}
		} else if info.Op == "open" {
			err := utils.SetOpenCourseSelect(info.Op)
			if err != nil {
				utils.Response(&ret, &w, err.Error())
				return
			}
		} else if info.Op == "close" {
			err := utils.SetOpenCourseSelect(info.Op)
			if err != nil {
				utils.Response(&ret, &w, err.Error())
				return
			}
		} else if info.Op == "queryOpen" {
			ret["open"], err = utils.IfOpenSelectCourse()
			if err != nil {
				utils.Response(&ret, &w, err.Error())
				return
			}
		} else {
			utils.Response(&ret, &w, "fail")
			return
		}
	}
	utils.Response(&ret, &w, "ok")
}
