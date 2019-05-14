package admin

import (
	"net/http"
)

func CourseManagement(w http.ResponseWriter, r *http.Request) {
	/*
		ret := make(map[string]interface{})

		// --- token 检查
		claims, err := utils.PreCheck(r)
		if err != nil {
			utils.Response(&ret, &w, err.Error())
			return
		}
		aid := claims["id"].(string)

		if r.Method == "GET" {
			ret["term"], err = utils.QueryTerm()
			if err != nil {
				utils.Response(&ret, &w, err.Error())
				return
			}
		} else if r.Method == "POST" {
			type Info struct {
				Op     string `json:"op"`
				Did    string `json:"did"`
				Tid    string `json:"tid"`
				Cid    string `json:"cid"`
				Credit string `json:"credit"`
				Term   string `json:"term"`
				Cname  string `json:"cname"`
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

			if info.Op == "select" {
				c, err := utils.QueryCourseWithTerm(info.Term)
				if err != nil {
					utils.Response(&ret, &w, err.Error())
					return
				}
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
		utils.Response(&ret, &w, "ok")
	*/
}
