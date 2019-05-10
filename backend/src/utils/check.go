package utils

import (
	"fmt"
	"log"
	"os"
	"strings"
	"taas/models"
)

func CheckId(id string) string {
	err := Db.QueryRow("select id from User where id = ?", id).Scan(&id)
	if err != nil {
		return "id not found"
	}
	return id
}
func CheckDname(dname string) string {
	var did string
	err := Db.QueryRow("select did from Department where dname = ?", dname).Scan(&did)
	if err != nil {
		return "dname not found"
	}
	return did
}

func CheckUser(id string, pwd string) (msg, t string) {
	t = "invaild"
	var retpwd string
	err := Db.QueryRow("select password,identity from User where id = ? ", id).Scan(&retpwd, &t)
	if err != nil {
		if err.Error() == models.NoRows {
			msg = "User not exits" // 用户不存在
			return
		}
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		msg = "fail"
		return
	}
	if strings.Compare(retpwd, pwd) != 0 {
		msg = "wrong password" // 密码错误
		return
	}
	msg = "ok"
	return
}

// true: 有联系
func CheckConnection(tid string) bool {
	var cnt int
	err := Db.QueryRow("select count(id) from CourseCalendar where tid = ?", tid).Scan(&cnt)
	if err != nil {
		log.Fatal(err)
	}
	return cnt != 0
}
