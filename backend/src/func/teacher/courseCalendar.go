package teacher

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
	tid := claims["id"].(string)

	if r.Method == "GET" {
		ret["term"], err = utils.GetAllTerms()
		if err != nil {
			utils.Response(&ret, &w, err.Error())
			return
		}
	} else if r.Method == "POST" {
		type Info struct {
			Term string `json:"term"`
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

		c, err := utils.QueryTeaCourse(tid, info.Term)
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
