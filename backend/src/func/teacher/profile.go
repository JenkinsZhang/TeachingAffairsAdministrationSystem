package teacher

import (
	"net/http"
	"taas/utils"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method != "GET" || len(r.Header["Authorization"]) == 0 {
		w.WriteHeader(http.StatusOK)
		return
	}
	ret := map[string]interface{}{
		"message": "ok",
	}
	var res = make(map[string]string)
	var id string
	token := r.Header["Authorization"][0]
	claims, err := utils.CheckToken(token)
	if err != nil {
		ret["message"] = "invalid token"
		utils.Response(ret, w)
		return
	}

	id = claims["id"].(string)
	res = utils.GetTeacherProfile(id)
	for key, val := range res {
		ret[key] = val
	}
	utils.Response(ret, w)
}
