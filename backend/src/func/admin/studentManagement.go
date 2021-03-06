package admin

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"taas/utils"
)

func StudentManagement(w http.ResponseWriter, r *http.Request) {
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
			Op         string `json:"op"`
			Id         string `json:"id"`
			Dname      string `json:"dname"`
			Did        string `json:"did"`
			Name       string `json:"name"`
			Gender     string `json:"gender"`
			Birthday   string `json:"birthday"`
			Birthplace string `json:"birthplace"`
			Phone      string `json:"phone"`
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
			err = utils.InsertStudent(utils.Struct2Map(info))
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
			ok, err := utils.CheckStuConnection(info.Id)
			if err != nil {
				utils.Response(&ret, &w, err.Error())
				return
			}
			if ok {
				utils.Response(&ret, &w, "this student have selected classes, so you can't delete it.")
				return
			}
			err = utils.DeleteStudent(info.Id)
			if err != nil {
				utils.Response(&ret, &w, err.Error())
				return
			}
		} else {
			utils.Response(&ret, &w, "invalid op")
			return
		}
	} else if r.Method == "GET" {
		c, err := utils.GetAllStudentProfile()
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
