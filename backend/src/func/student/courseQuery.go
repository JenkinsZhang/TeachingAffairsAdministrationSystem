package student

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"taas/models"
	"taas/utils"
)

func CourseQuery(w http.ResponseWriter, r *http.Request) {
	ret := make(map[string]interface{})

	// --- token 检查
	claims, err := utils.PreCheck(r)
	if err != nil {
		utils.Response(&ret, &w, err.Error())
		return
	}
	id := claims["id"].(string)

	// ---

	if r.Method == "POST" {
		type Info struct {
			Cid       string `json:"cid"`
			Tid       string `json:"tid"`
			Cname     string `json:"cname"`
			Tname     string `json:"tname"`
			Term      string `json:"term"`
			ClassTime string `json:"ClassTime"`
			Op        string `json:"op"`
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

		c := make(map[string][]string)

		if info.Op == "query" {
			if info.Cid != "" {
				c, err = utils.QueryCourseWithCid(info.Cid)
			} else if info.Cname != "" {
				c, err = utils.QueryCourseWithCname(info.Cname)
			} else if info.Tid != "" {
				c, err = utils.QueryCourseWithTid(info.Tid)
			} else if info.Tname != "" {
				c, err = utils.QueryCourseWithTname(info.Tname)
			} else {
				c, err = utils.QueryAllCourses()
			}
			if err != nil {
				utils.Response(&ret, &w, err.Error())
				return
			}
			for key, val := range c {
				ret[key] = val
			}
		} else if info.Op == "select" {
			if models.OpenSelectCourse == 0 {
				utils.Response(&ret, &w, "Lesson selection time is not yet available.")
				return
			}
			tmp := utils.Struct2Map(info)
			tmp["term"], err = utils.GetCurrentTerm()
			if err != nil {
				utils.Response(&ret, &w, err.Error())
				return
			}
			tmp["Id"] = id
			err = utils.InsertCourse(tmp)
			if err != nil {
				utils.Response(&ret, &w, err.Error())
				return
			}
		}
	}
	utils.Response(&ret, &w, "ok")
}
