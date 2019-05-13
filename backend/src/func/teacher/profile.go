package teacher

import (
	"net/http"
	"taas/utils"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	ret := make(map[string]interface{})

	// --- token 检查
	claims, err := utils.PreCheck(r)
	if err != nil {
		utils.Response(&ret, &w, err.Error())
		return
	}
	tid := claims["id"].(string)

	if r.Method == "GET" {
		res, err := utils.GetTeacherProfile(tid)
		if err != nil {
			utils.Response(&ret, &w, err.Error())
			return
		}
		for key, val := range res {
			ret[key] = val
		}
	}
	utils.Response(&ret, &w, "ok")
}
