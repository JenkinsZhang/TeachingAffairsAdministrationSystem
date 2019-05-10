package teacher

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"taas/utils"
)

func CourseCalendar(w http.ResponseWriter, r *http.Request) {
	type Info struct {
		Term string `json:"term"`
	}
	r.ParseForm()
	arr, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "readall: %v\n", err)
		os.Exit(-1)
	}
	var info Info
	err = json.Unmarshal(arr, &info)
	if err != nil {
		fmt.Fprintf(os.Stderr, "unmarshal: %v\n", err)
		os.Exit(-1)
	}
	if r.Method != "POST" || len(r.Header["Authorization"]) == 0 {
		w.WriteHeader(http.StatusOK)
		return
	}

	ret := make(map[string]interface{})
	ret["message"] = "ok"
	token := r.Header["Authorization"][0]
	claims, err := utils.CheckToken(token)
	tid := claims["id"].(string)
	if err != nil {
		ret["message"] = "invalid token"
		utils.Response(ret, w)
		return
	}
	// ----
	c := utils.QueryTeaCourse(tid, info.Term)
	for key, val := range c {
		ret[key] = val
	}
	utils.Response(ret, w)
}
