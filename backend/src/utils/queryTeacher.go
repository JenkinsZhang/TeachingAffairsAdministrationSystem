package utils

import (
	"fmt"
	"log"
	"os"
)

func QueryTeaDepartmentId(tid string) string {
	var did string
	err := Db.QueryRow("select did from Teacher where tid = ? ", tid).Scan(&did)
	if err != nil {
		log.Fatal(err)
	}
	return did
}
func GetTeacherProfile(tid string) map[string]string {
	var tname, gender, did, dname string
	err := Db.QueryRow("select tname, gender, did from Teacher where tid = ?", tid).Scan(&tname, &gender, &did)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	dname = QueryDepartmentName(did)
	ret := map[string]string{
		"tname":  tname,
		"gender": gender,
		"dname":  dname,
		"tid":    tid,
	}
	return ret
}

func QueryStuWithTidCid(tid, cid, term string) map[string][]string {
	var id, did, dname, score string
	rows, err := Db.Query("select id, score from CourseCalendar where tid = ? and cid = ? and term = ?", tid, cid, term)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	ret := make(map[string][]string)

	for rows.Next() {
		err := rows.Scan(&id, &score)
		if err != nil {
			log.Fatal(err)
		}
		did = QueryStuDepartmentId(id)
		dname = QueryDepartmentName(did)
		name := QueryStuName(id)
		ret["id"] = append(ret["id"], id)
		ret["name"] = append(ret["name"], name)
		ret["dname"] = append(ret["dname"], dname)
		ret["score"] = append(ret["score"], score)
	}
	return ret
}

func QueryCourseScore(tid, cid, term string) []string {
	rows, err := Db.Query("select score from CourseCalendar where tid = ? and cid = ? and term = ?", tid, cid, term)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	ret := make([]string, 0)
	var score string
	for rows.Next() {
		err := rows.Scan(&score)
		if err != nil {
			log.Fatal(err)
		}
		ret = append(ret, score)
	}
	return ret
}

func QueryTeaCourse(tid, term string) map[string][]string {
	var cid, classTime, cname, credit string
	rows, err := Db.Query("select cid,classTime from CourseCalendar where tid = ? and term = ?", tid, term)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	ret := make(map[string][]string)

	for rows.Next() {
		err := rows.Scan(&cid, &classTime)
		if err != nil {
			log.Fatal(err)
		}
		cname = QueryCourseName(cid)
		credit = QueryCourseCredit(cid)
		ret["cid"] = append(ret["cid"], cid)
		ret["cname"] = append(ret["cname"], cname)
		ret["credit"] = append(ret["credit"], credit)
		ret["classTime"] = append(ret["classTime"], classTime)
	}
	return ret
}

func QueryCourseStuInfo(tid, cid, term string) map[string][]string {
	var id string
	rows, err := Db.Query("select id from CourseCalendar where tid = ? and cid = ? and term = ?", tid, cid, term)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	ret := make(map[string][]string, 0)
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
		c := GetStudentProfile(id)
		for key, val := range c {
			ret[key] = append(ret[key], val)
		}
	}
	return ret
}
