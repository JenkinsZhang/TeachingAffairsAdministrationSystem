package admin

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
	token := r.Header["Authorization"][0]
	claims, err := utils.CheckToken(token)
	if err != nil {
		ret["message"] = "invalid token"
		utils.Response(ret, w)
		return
	}

	aid := claims["id"].(string)
	ret["aid"] = aid
	utils.Response(ret, w)

}
