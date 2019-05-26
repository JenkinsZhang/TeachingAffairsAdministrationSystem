package utils

import (
	"errors"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func CheckTidInTeacher(tid string) error {
	err := Db.QueryRow("select tid from Teacher where tid = ?", tid).Scan(&tid)
	return err
}
func CheckCidInCourse(cid string) error {
	err := Db.QueryRow("select cid from Course where cid = ?", cid).Scan(&cid)
	return err
}
func CheckCnameInCourse(cname string) error {
	err := Db.QueryRow("select cname from Course where cname = ?", cname).Scan(&cname)
	return err
}

func PreCheck(r *http.Request) (jwt.MapClaims, error) {
	r.ParseForm()
	if len(r.Header["Authorization"]) == 0 {
		return nil, errors.New("need token")
	}
	token := r.Header["Authorization"][0]
	claims, err := CheckToken(token)
	if err != nil {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
func CheckId(id string) error {
	err := Db.QueryRow("select id from User where id = ?", id).Scan(&id)
	return err
}
func CheckDname(dname string) (string, error) {
	var did string
	err := Db.QueryRow("select did from Department where dname = ?", dname).Scan(&did)
	return did, err
}

// true: 有联系
func CheckTeaConnection(tid string) (bool, error) {
	var cnt int
	err := Db.QueryRow("select count(id) from CourseCalendar where tid = ?", tid).Scan(&cnt)
	if err != nil {
		return false, err
	}
	return cnt != 0, nil
}
func CheckStuConnection(id string) (bool, error) {
	var cnt int
	err := Db.QueryRow("select count(cid) from CourseCalendar where id = ?", id).Scan(&cnt)
	if err != nil {
		return false, err
	}
	return cnt != 0, nil
}

func CheckUser(id string, pwd string) (string, error) {
	t := "invaild"
	var retpwd string
	err := Db.QueryRow("select password,identity from User where id = ? ", id).Scan(&retpwd, &t)
	if err != nil {
		return t, err
	}
	if strings.Compare(retpwd, pwd) != 0 {
		return t, errors.New("wrong password")
	}
	return t, nil
}

// ----------------------------

/*
func GetJson(r *http.Request, info interface{}) interface{}, error) {
	arr, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(arr, &info)
	if err != nil {
		return nil, err
	}
	return info, nil
}
*/
