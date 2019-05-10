package teacher

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"taas/utils"
)

func ScoreManagement(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if len(r.Header["Authorization"]) == 0 {
		w.WriteHeader(http.StatusOK)
		return
	}
	ret := make(map[string]interface{})

	token := r.Header["Authorization"][0]
	claims, err := utils.CheckToken(token)
	if err != nil {
		ret["message"] = "invalid token"
		utils.Response(ret, w)
		return
	}
	tid := claims["id"].(string)
	if r.Method == "GET" {
		c := utils.QueryCourseWithTid(tid)
		ret["term"] = c["term"]
		ret["cid"] = c["cid"]
		ret["cname"] = c["cname"]
		utils.Response(ret, w)
	} else if r.Method == "POST" {
		type Info struct {
			Cid   string `json:"cid"`
			Id    string `json:"id"`
			Score string `json:"score"`
			Term  string `json:"term"`
			Op    string `json:"op"`
		}
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
		// ----
		if info.Op == "query" {
			c := utils.QueryStuWithTidCid(tid, info.Cid, info.Term)
			for key, val := range c {
				ret[key] = val
			}
		} else if info.Op == "change" {
			ret["message"] = utils.UpdateStuScore(info.Id, info.Cid, info.Score, info.Term)
		} else {
			ret["message"] = "fail"
		}
		utils.Response(ret, w)
	}
}
