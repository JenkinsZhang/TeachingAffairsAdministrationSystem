package login

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"taas/utils"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

/*
func needToken(handleFunc func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request){
	return func (w http.ResponseWriter, r *http.Request)  {
		CheckUser()
		handleFunc(w,r)
	}
}
*/

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	arr, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "readall: %v\n", err)
		os.Exit(-1)
	}
	if r.Method == "POST" {
		var user User
		err = json.Unmarshal(arr, &user)
		if err != nil {
			fmt.Fprintf(os.Stderr, "unmarshal: %v\n", err)
			os.Exit(-1)
		}
		msg, identity := utils.CheckUser(user.Id, user.Password)
		name := "admin"
		if identity == "student"{
			name = utils.QueryStuName(user.Id)
		} else if identity == "teacher"{
			name = utils.QueryTeaName(user.Id)
		}
		payload := jwt.MapClaims{
			"name":		name,
			"id":       user.Id,
			"identity": identity,
		}
		tmp := map[string]interface{}{
			"message":  msg,
			"identity": identity,
			"token":    utils.NewToken(payload),
		}
		jtmp, err := json.Marshal(tmp)
		if err != nil {
			fmt.Fprintf(os.Stderr, "marshal: %v\n", err)
			tmp["message"] = "fail"
		}
		w.Write(jtmp)
	}
	w.WriteHeader(http.StatusOK)
}

/*
id:1621674
token:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjE2MTIxNjc0IiwiaWRlbnRpdHkiOiJzdHVkZW50In0.WUdU2mE6Kd3ex9ZXrOBaQgkHAMUwzpSdgInUh6Grt-M"

id:10001001
tokne:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEwMDAxMDAxIiwiaWRlbnRpdHkiOiJ0ZWFjaGVyIn0.BN8j7W_x6Sosp1SsFCL4kqBhTYYCP02Ialpz3-rlBvM

id:9999
token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6Ijk5OTkiLCJpZGVudGl0eSI6ImFkbWluIn0.rteWPgRmbegYeuI9gwnhl3GWc5F6-i2j8q3ZL1_sWWY
*/
