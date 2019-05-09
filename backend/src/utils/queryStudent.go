package utils

import (
	"fmt"
	"log"
	"os"
)

func GetStudentProfile(id string) map[string]string {

	var name, gender, did, dname, grade, phone, birthplace string
	err := Db.QueryRow("select name, gender, did, grade,phone,birthplace from Student where id = ?", id).Scan(&name, &gender, &did, &grade, &phone, &birthplace)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	dname = QueryDepartmentName(did)
	ret := map[string]string{
		"name":       name,
		"gender":     gender,
		"dname":      dname,
		"grade":      grade,
		"id":         id,
		"phone":      phone,
		"birthplace": birthplace,
	}
	return ret
}

func QueryStuCourses(id, term string) map[string][]string {
	ret := make(map[string][]string)
	rows, err := Db.Query("select CourseCalendar.cid, cname,CourseCalendar.tid,tname, credit, classTime from CourseCalendar, Course, Teacher where CourseCalendar.id = ? and CourseCalendar.cid = Course.cid and CourseCalendar.tid = Teacher.tid and term = ?", id, term)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer rows.Close()
	var cid, tid, tname, cname, classTime, credit string
	for rows.Next() {
		err := rows.Scan(&cid, &cname, &tid, &tname, &credit, &classTime)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		ret["cid"] = append(ret["cid"], cid)
		ret["tid"] = append(ret["tid"], tid)
		ret["tname"] = append(ret["tname"], tname)
		ret["cname"] = append(ret["cname"], cname)
		ret["classTime"] = append(ret["classTime"], classTime)
		ret["credit"] = append(ret["credit"], credit)
	}
	return ret
}

func QueryAllCourses() map[string][]string {
	ret := make(map[string][]string)
	rows, err := Db.Query("select Course.cid, cname, Teacher.tid, tname, credit, classTime from Course, CourseSchedule, Teacher where CourseSchedule.cid = Course.cid and CourseSchedule.tid = Teacher.tid")

	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer rows.Close()
	var tid, tname, cid, cname, classTime, credit string
	for rows.Next() {
		err := rows.Scan(&cid, &cname, &tid, &tname, &credit, &classTime)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		ret["cid"] = append(ret["cid"], cid)
		ret["tid"] = append(ret["tid"], tid)
		ret["tname"] = append(ret["tname"], tname)
		ret["cname"] = append(ret["cname"], cname)
		ret["classTime"] = append(ret["classTime"], classTime)
		ret["credit"] = append(ret["credit"], credit)
	}
	return ret
}

func QueryTotalStu(did string) string {
	rows, err := Db.Query("select count(id) from Student where did = ?", did)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var cnt int
	for rows.Next() {
		err := rows.Scan(&cnt)
		if err != nil {
			log.Fatal(err)
		}
	}
	return fmt.Sprintf("%d", cnt)
}

func QueryStuName(id string) string {
	var name string
	err := Db.QueryRow("select name from Student where id = ?", id).Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	return name
}
