package admin

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"taas/utils"
)

func TeacherManagement(w http.ResponseWriter, r *http.Request) {
	ret := make(map[string]interface{})

	// --- token 检查
	_, err := utils.PreCheck(r)
	if err != nil {
		utils.Response(&ret, &w, err.Error())
		return
	}
	// aid := claims["id"].(string)

	// ----
	if r.Method == "POST" {
		type Info struct {
			Op        string `json:"op"`
			Dname     string `json:"dname"`
			Tid       string `json:"tid"`
			Did       string `json:"did"`
			Tname     string `json:"tname"`
			Gender    string `json:"gender"`
			Birthday  string `json:"birthday"`
			Education string `json:"education"`
			Wage      string `json:"wage"`
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

		if info.Op == "add" {
			did, err := utils.CheckDname(info.Dname)
			if err != nil {
				utils.Response(&ret, &w, err.Error())
				return
			}
			info.Did = did
			err = utils.InsertTeacher(utils.Struct2Map(info))
			if err != nil {
				utils.Response(&ret, &w, err.Error())
				return
			}
		} else if info.Op == "modify" {
			err = utils.Update(utils.Struct2Map(info))
			if err != nil {
				utils.Response(&ret, &w, err.Error())
				return
			}
		} else if info.Op == "delete" {
			ok, err := utils.CheckConnection(info.Tid)
			if err != nil {
				utils.Response(&ret, &w, "still have lessons")
				return
			}
			if ok {
				utils.Response(&ret, &w, "fail")
				return
			}
			err = utils.DeleteTeacher(info.Tid)
			if err != nil {
				utils.Response(&ret, &w, err.Error())
				return
			}

		} else {
			utils.Response(&ret, &w, "invalid op")
			return
		}
	} else if r.Method == "GET" {
		c, err := utils.GetAllTeacherProfile()
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
