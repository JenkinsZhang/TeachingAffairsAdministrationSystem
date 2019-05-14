package student

import (
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
		term, err := utils.QueryTerm()
		if err != nil {
			utils.Response(&ret, &w, err.Error())
			return
		}
		score := make([]string, 0)
		for _, v := range term {
			tmp, err := utils.QueryTermAveScore(id, v)
			if err != nil {
				utils.Response(&ret, &w, err.Error())
				return
			}
			score = append(score, tmp)
		}
		ret["term"] = term
		ret["score"] = score
	}
	utils.Response(&ret, &w, "ok")
}
