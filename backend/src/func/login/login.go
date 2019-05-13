package login

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
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
	ret := make(map[string]interface{})
	ret["message"] = "ok"
	if r.Method == "POST" {
		arr, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			ret["message"] = err.Error()
			utils.Response(ret, w)
			return
		}
		var user User
		err = json.Unmarshal(arr, &user)
		if err != nil {
			ret["message"] = err.Error()
			utils.Response(ret, w)
			return
		}
		msg, identity := utils.CheckUser(user.Id, user.Password)
		name := "admin"
		if identity == "student" {
			name, err = utils.QueryStuName(user.Id)
		} else if identity == "teacher" {
			name, err = utils.QueryTeaName(user.Id)
		}
		if err != nil {
			ret["message"] = err.Error()
			utils.Response(ret, w)
			return
		}
		payload := jwt.MapClaims{
			"name":     name,
			"id":       user.Id,
			"identity": identity,
		}
		ret["message"] = msg
		ret["identity"] = identity
		ret["token"], err = utils.NewToken(payload)
		if err != nil {
			ret["message"] = err.Error()
			utils.Response(ret, w)
			return
		}
	}
	utils.Response(ret, w)
}

/*
id:1621674
token:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjE2MTIxNjc0IiwiaWRlbnRpdHkiOiJzdHVkZW50In0.WUdU2mE6Kd3ex9ZXrOBaQgkHAMUwzpSdgInUh6Grt-M"

id:10001001
tokne:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjEwMDAxMDAxIiwiaWRlbnRpdHkiOiJ0ZWFjaGVyIn0.BN8j7W_x6Sosp1SsFCL4kqBhTYYCP02Ialpz3-rlBvM

id:9999
token:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6Ijk5OTkiLCJpZGVudGl0eSI6ImFkbWluIn0.rteWPgRmbegYeuI9gwnhl3GWc5F6-i2j8q3ZL1_sWWY
*/
