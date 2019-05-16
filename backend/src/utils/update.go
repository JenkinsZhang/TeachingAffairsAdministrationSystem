package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func UpdateStuScore(id, cid, score, term string) error {
	stmt, err := Db.Prepare("update CourseCalendar set score = ? where id = ? and cid = ? and term = ?")
	if err != nil {
		return err
	}
	tmp, err := strconv.ParseFloat(score, 32)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(tmp, id, cid, term)
	return err
}

func UpdateInfo(id, itemName, itemVal, table, idName string) error {
	q := fmt.Sprintf("update %s set %s = ? where %s = ?", table, itemName, idName)
	stmt, err := Db.Prepare(q)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(itemVal, id)
	return err
}
func Update(info map[string]string) error {
	if _, ok := info["Id"]; ok {
		items := []string{"Name", "Gender", "Birthday", "Birthplace", "Phone"}
		id := info["Id"]
		stu := "Student"
		for _, itemName := range items {
			if itemVal, ok := info[itemName]; ok && len(itemVal) != 0 {
				err := UpdateInfo(id, strings.ToLower(itemName), itemVal, stu, "id")
				if err != nil {
					return err
				}
			}
		}
		if dname, ok := info["Dname"]; ok && len(dname) != 0 {
			did, err := CheckDname(dname)
			if err != nil {
				return err
			}
			err = UpdateInfo(id, "did", did, stu, "id")
			if err != nil {
				return err
			}
		}
	} else if _, ok := info["Tid"]; ok {
		items := []string{"Tname", "Gender", "Birthday", "Education", "Wage"}
		tid := info["Tid"]
		tea := "Teacher"
		for _, itemName := range items {
			if itemVal, ok := info[itemName]; ok && len(itemVal) != 0 {
				err := UpdateInfo(tid, strings.ToLower(itemName), itemVal, tea, "tid")
				if err != nil {
					return err
				}
			}
		}
		if dname, ok := info["Dname"]; ok && len(dname) != 0 {
			did, err := CheckDname(dname)
			if err != nil {
				return err
			}
			err = UpdateInfo(tid, "did", did, tea, "tid")
			if err != nil {
				return err
			}
		}
	}
	return nil
}
