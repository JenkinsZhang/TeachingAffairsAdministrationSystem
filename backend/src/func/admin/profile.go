package admin

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
	aid := claims["id"].(string)

	if r.Method == "GET" {
		ret["aid"] = aid
	}
	utils.Response(&ret, &w, "ok")
}
