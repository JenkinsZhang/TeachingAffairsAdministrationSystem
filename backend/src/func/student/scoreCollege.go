package student

import (
	"fmt"
	"net/http"
	"strconv"
	"taas/utils"
)

func ScoreCollege(w http.ResponseWriter, r *http.Request) {
	ret := make(map[string]interface{})

	// --- token 检查
	claims, err := utils.PreCheck(r)
	if err != nil {
		utils.Response(&ret, &w, err.Error())
		return
	}
	id := claims["id"].(string)

	
	// ----
	if r.Method == "GET"{
		did,err := utils.QueryStuDepartmentId(id)
		if err != nil {
			utils.Response(&ret, &w, err.Error())
			return
		}
		totalStudents,err := utils.QueryTotalStu(did)
		if err != nil {
			utils.Response(&ret, &w, err.Error())
			return
		}
		score,err := utils.QueryStuAveScore(id)
		if err != nil {
			utils.Response(&ret, &w, err.Error())
			return
		}
		rank,err := utils.QueryStuRank(id, did)
		if err != nil {
			utils.Response(&ret, &w, err.Error())
			return
		}

		t1, _ := strconv.ParseFloat(rank, 32)
		t2, _ := strconv.ParseFloat(totalStudents, 32)
		percentage := t1 / t2

		ret["id"] = id
		ret["did"],err = utils.QueryDepartmentName(did)
		if err != nil {
			utils.Response(&ret, &w, err.Error())
			return
		}
		ret["score"] = score
		ret["totalStudents"] = totalStudents
		ret["rank"] = rank
		ret["percentage"] = fmt.Sprintf("%f", percentage)
		ret["name"],err = utils.QueryStuName(id)
		if err != nil {
			utils.Response(&ret, &w, err.Error())
			return
		}
	}
	utils.Response(&ret, &w, "ok")
}
