package student

import (
	"net/http"
	"taas/utils"
)

func ScoreSummary(w http.ResponseWriter, r *http.Request) {
	ret := make(map[string]interface{})

	// --- token 检查
	claims, err := utils.PreCheck(r)
	if err != nil {
		utils.Response(&ret, &w, err.Error())
		return
	}
	id := claims["id"].(string)

	// ----
	if r.Method == "GET" {
		c, err := utils.QueryStuAllCourseScore(id)
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
