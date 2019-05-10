package student

import (
	"net/http"
	"taas/utils"
)

func ScoreSummary(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method != "GET" || len(r.Header["Authorization"]) == 0 {
		w.WriteHeader(http.StatusOK)
		return
	}

	ret := make(map[string]interface{})
	ret["message"] = "ok"
	token := r.Header["Authorization"][0]
	claims, err := utils.CheckToken(token)
	id := claims["id"].(string)
	if err != nil {
		ret["message"] = "invalid token"
		utils.Response(ret, w)
		return
	}
	// ----
	c := utils.QueryStuAllCourseScore(id)
	for key, val := range c {
		ret[key] = val
	}

	utils.Response(ret, w)
}
