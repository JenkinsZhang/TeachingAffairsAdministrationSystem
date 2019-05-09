package utils

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func UpdateStuScore(id, cid, score, term string) string {
	stmt, err := Db.Prepare("update CourseCalendar set score = ? where id = ? and cid = ? and term = ?")
	if err != nil {
		log.Println(err)
		return "fail"
	}
	tmp, _ := strconv.ParseFloat(score, 32)
	_, err = stmt.Exec(tmp, id, cid, term)
	if err != nil {
		log.Println(err)
		return "fail"
	}
	return "ok"
}

func UpdateInfo(id, itemName, itemVal, table, idName string) {
	q := fmt.Sprintf("update %s set %s = ? where %s = ?", table, itemName, idName)
	stmt, err := Db.Prepare(q)
	if err != nil {
		log.Fatal(err)
	}
	_, err = stmt.Exec(itemVal, id)
	if err != nil {
		log.Fatal(err)
	}
}
func Update(info map[string]string) string {
	if _, ok := info["Id"]; ok {
		items := []string{"Name", "Gender", "Birthday", "Birthplace", "Phone", "Grade"}
		id := info["Id"]
		stu := "Student"
		for _, itemName := range items {
			if itemVal, ok := info[itemName]; ok && len(itemVal) != 0 {
				UpdateInfo(id, strings.ToLower(itemName), itemVal, stu, "id")
			}
		}
		if dname, ok := info["Dname"]; ok && len(dname) != 0 {
			did := CheckDname(dname)
			if did != "dname not found" {
				UpdateInfo(id, "did", did, stu, "id")
			} else {
				return "fail"
			}
		}
	} else if _, ok := info["Tid"]; ok {
		items := []string{"Tname", "Gender", "Birthday", "Education", "Wage"}
		tid := info["Tid"]
		tea := "Teacher"
		for _, itemName := range items {
			if itemVal, ok := info[itemName]; ok && len(itemVal) != 0 {
				UpdateInfo(tid, strings.ToLower(itemName), itemVal, tea, "tid")
			}
		}
		if dname, ok := info["Dname"]; ok && len(dname) != 0 {
			did := CheckDname(dname)
			if did != "dname not found" {
				UpdateInfo(tid, "did", did, tea, "tid")
			} else {
				return "fail"
			}
		}
	}
	return "ok"
}
