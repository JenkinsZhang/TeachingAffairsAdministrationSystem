package teacher

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"taas/utils"
)

func ClassTable(w http.ResponseWriter, r *http.Request) {
	ret := make(map[string]interface{})

	// --- token 检查
	claims, err := utils.PreCheck(r)
	if err != nil {
		utils.Response(&ret, &w, err.Error())
		return
	}
	tid := claims["id"].(string)

	if r.Method == "GET" {
		c, err := utils.QueryCourseWithTid(tid)
		if err != nil {
			utils.Response(&ret, &w, err.Error())
			return
		}
		ret["term"] = c["term"]
		ret["cid"] = c["cid"]
		ret["cname"] = c["cname"]
	} else if r.Method == "POST" {
		type Info struct {
			Term string `json:"term"`
			Cid  string `json:"cid"`
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

		c, err := utils.QueryCourseStuInfo(tid, info.Cid, info.Term)
		if err != nil {
			utils.Response(&ret, &w, err.Error())
			return
		}
		for key, val := range c {
			ret[key] = val
		}
	}
	utils.Response(&ret, &w, "ok")
}
