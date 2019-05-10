package admin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"taas/utils"
)

func StudentManagement(w http.ResponseWriter, r *http.Request) {
	type Info struct {
		Op         string `json:"op"`
		Id         string `json:"id"`
		Dname      string `json:"dname"`
		Did        string `json:"did"`
		Name       string `json:"name"`
		Gender     string `json:"gender"`
		Birthday   string `json:"birthday"`
		Birthplace string `json:"birthplace"`
		Grade      string `json:"grade"`
		Phone      string `json:"phone"`
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
	_, err = utils.CheckToken(token)
	if err != nil {
		ret["message"] = "invalid token"
		utils.Response(ret, w)
		return
	}
	// ----
	if info.Op == "add" {
		did := utils.CheckDname(info.Dname)
		if did == "dname not found" {
			ret["message"] = "fail"
			utils.Response(ret, w)
			return
		}
		if utils.CheckId(info.Id) != "id not found" {
			ret["message"] = "fail"
			utils.Response(ret, w)
			return
		}
		info.Did = did
		utils.InsertStudent(utils.Struct2Map(info))
	} else if info.Op == "modify" {
		ret["message"] = utils.Update(utils.Struct2Map(info))
	} else if info.Op == "delete" {
		ok := utils.CheckConnection(info.Id)
		if ok {
			ret["message"] = "fail"
			utils.Response(ret, w)
			return
		}
		utils.DeleteStudent(info.Id)
	} else {
		ret["message"] = "fail"
	}
	utils.Response(ret, w)
}
