package student

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"taas/utils"
)

func ScoreTrend(w http.ResponseWriter, r *http.Request) {
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
		type Info struct {
			Term []string `json:"term"`
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

		score := make([]string, 0)
		for _, v := range info.Term {
			tmp, err := utils.QueryTermAveScore(id, v)
			if err != nil {
				utils.Response(&ret, &w, err.Error())
				return
			}
			score = append(score, tmp)
		}
		ret["score"] = score
	}
	utils.Response(&ret, &w, "ok")
}
