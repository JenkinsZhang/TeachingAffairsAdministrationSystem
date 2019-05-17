package teacher

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"taas/utils"
)

func ScoreManagement(w http.ResponseWriter, r *http.Request) {
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
			fmt.Println(err)
			return
		}
		ret["term"] = c["term"]
		ret["cid"] = c["cid"]
		ret["cname"] = c["cname"]
	} else if r.Method == "POST" {
		type Info struct {
			Cid   string `json:"cid"`
			Id    string `json:"id"`
			Score string `json:"score"`
			Term  string `json:"term"`
			Op    string `json:"op"`
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

		if info.Op == "query" {
			c, err := utils.QueryStuWithTidCid(tid, info.Cid, info.Term)
			if err != nil {
				utils.Response(&ret, &w, err.Error())
				return
			}
			for key, val := range c {
				ret[key] = val
			}
		} else if info.Op == "change" {
			term, err := utils.GetCurrentTerm()
			if err != nil {
				utils.Response(&ret, &w, err.Error())
				return
			}
			if term != info.Term { // 不是当前学期
				utils.Response(&ret, &w, "not current term")
				return
			}
			err = utils.UpdateStuScore(info.Id, info.Cid, info.Score, info.Term)
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
