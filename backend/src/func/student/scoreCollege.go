package student

import (
	"fmt"
	"net/http"
	"strconv"
	"taas/utils"
)

func ScoreCollege(w http.ResponseWriter, r *http.Request) {
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
	did := utils.QueryStuDepartmentId(id)
	totalStudents := utils.QueryTotalStu(did)
	score := utils.QueryStuAveScore(id)
	rank := utils.QueryStuRank(id, did)

	t1, _ := strconv.ParseFloat(rank, 32)
	t2, _ := strconv.ParseFloat(totalStudents, 32)
	percentage := t1 / t2

	ret["id"] = id
	ret["did"] = utils.QueryDepartmentName(did)
	ret["score"] = score
	ret["totalStudents"] = totalStudents
	ret["rank"] = rank
	ret["percentage"] = fmt.Sprintf("%f", percentage)
	ret["name"] = utils.QueryStuName(id)

	utils.Response(ret, w)

}
