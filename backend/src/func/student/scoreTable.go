package student

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"taas/utils"
)

func ScoreTable(w http.ResponseWriter, r *http.Request) {
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
	if  len(r.Header["Authorization"]) == 0 {
		w.WriteHeader(http.StatusOK)
		return
	}

	ret := make(map[string]interface{})
	ret["message"] = "ok"
	if r.Method == "GET"{
		ret["term"] = utils.QueryTerm()
	} else if r.Method == "POST"{
		token := r.Header["Authorization"][0]
		claims, err := utils.CheckToken(token)
		id := claims["id"].(string)
		if err != nil {
			ret["message"] = "invalid token"
			utils.Response(ret, w)
			return
		}
		// ----
		c := utils.QueryStuCourseScore(id, info.Term)
		for key, val := range c {
			ret[key] = val
		}
	}
	utils.Response(ret, w)
}
