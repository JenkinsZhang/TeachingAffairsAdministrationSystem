package admin

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"taas/utils"
)

func TermManagement(w http.ResponseWriter, r *http.Request) {
	ret := make(map[string]interface{})

	// --- token 检查
	_, err := utils.PreCheck(r)
	if err != nil {
		utils.Response(&ret, &w, err.Error())
		return
	}

	if r.Method == "GET" {
		c, err := utils.GetAllTerms()
		if err != nil {
			utils.Response(&ret, &w, err.Error())
			return
		}
		for key, val := range c {
			ret[key] = val
		}
	} else if r.Method == "POST" {
		type Info struct {
			Op   string `json:"op"`
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

		if info.Op == "add" {
			err = utils.InsertTerm(info.Term)
			if err != nil {
				utils.Response(&ret, &w, err.Error())
				return
			}
		} else if info.Op == "delete" {
			err = utils.DeleteTerm(info.Term)
			if err != nil {
				utils.Response(&ret, &w, err.Error())
				return
			}
		} else if info.Op == "end" {
			err = utils.EndTerm(info.Term)
			if err != nil {
				utils.Response(&ret, &w, err.Error())
				return
			}
		} else if info.Op == "set" {
			err = utils.SetCurrentTerm(info.Term)
			if err != nil {
				utils.Response(&ret, &w, err.Error())
				return
			}
		} else {
			utils.Response(&ret, &w, "invalid op")
			return
		}
	}
	utils.Response(&ret, &w, "ok")
}
