package student

import (
	"net/http"
	"taas/utils"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	ret := make(map[string]interface{})

	// --- token 检查

	if err != nil {
		utils.Response(&ret, &w, err.Error())
		return
	}
	id := claims["id"].(string)

	if r.Method == "GET" {
		ret := map[string]interface{}{
			"id":     "-1",
			"name":   "-1",
			"gender": "-1",
			"dname":  "-1",
			"grade":  "-1",
		}
		var res = make(map[string]string)
		res, err = utils.GetStudentProfile(id)
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
