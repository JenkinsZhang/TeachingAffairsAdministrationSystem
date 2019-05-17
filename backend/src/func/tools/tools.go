package tools

import (
	"net/http"
	"taas/utils"
)

func GetDepartment(w http.ResponseWriter, r *http.Request) {
	ret := make(map[string]interface{})
	c, err := utils.GetDepartment()
	if err != nil {
		utils.Response(&ret, &w, err.Error())
		return
	}
	for key, val := range c {
		ret[key] = val
	}
	utils.Response(&ret, &w, "ok")
}

func GetCurrentTerm(w http.ResponseWriter, r *http.Request) {
	term, err := utils.GetCurrentTerm()
	ret := make(map[string]interface{})
	ret["term"] = term
	if err != nil {
		utils.Response(&ret, &w, err.Error())
		return
	}
	utils.Response(&ret, &w, "ok")
}
