package utils

import (
	"fmt"
	"log"
	"os"
	"taas/models"
)

func QueryDepartmentName(did string) string {
	var dname string
	err := Db.QueryRow("select dname from Department where did = ?", did).Scan(&dname)
	if err != nil {
		if err.Error() == models.NoRows {
			return models.NoRows
		}
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
	return dname
}
func QueryStuDepartmentId(id string) string {
	var did string
	err := Db.QueryRow("select did from Student where id = ? ", id).Scan(&did)
	if err != nil {
		log.Fatal(err)
	}
	return did
}
func QueryCourseWithCid(cid string) map[string][]string {
	ret := make(map[string][]string)
	rows, err := Db.Query("select Course.cid, cname, Teacher.tid, tname, credit, classTime from Course, CourseSchedule, Teacher where CourseSchedule.cid = Course.cid and CourseSchedule.tid = Teacher.tid and Course.cid = ?", cid)

	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer rows.Close()
	var tid, tname, cname, classTime, credit string
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
func QueryCourseWithTid(tid string) map[string][]string {
	ret := make(map[string][]string)
	rows, err := Db.Query("select Course.cid, cname, Teacher.tid, tname, credit, classTime,term from Course, CourseSchedule, Teacher where CourseSchedule.cid = Course.cid and CourseSchedule.tid = Teacher.tid and Teacher.tid = ?", tid)

	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer rows.Close()
	var cid, tname, cname, classTime, credit, term string
	for rows.Next() {
		err := rows.Scan(&cid, &cname, &tid, &tname, &credit, &classTime, &term)
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
		ret["term"] = append(ret["term"], term)
	}
	return ret
}
func QueryCourseWithCname(cname string) map[string][]string {
	ret := make(map[string][]string)
	rows, err := Db.Query("select Course.cid, cname, Teacher.tid, tname, credit, classTime from Course, CourseSchedule, Teacher where CourseSchedule.cid = Course.cid and CourseSchedule.tid = Teacher.tid and Course.cname = ?", cname)

	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer rows.Close()
	var cid, tid, tname, classTime, credit string
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
func QueryCourseWithTname(tname string) map[string][]string {
	ret := make(map[string][]string)
	rows, err := Db.Query("select Course.cid, cname, Teacher.tid, tname, credit, classTime from Course, CourseSchedule, Teacher where CourseSchedule.cid = Course.cid and CourseSchedule.tid = Teacher.tid and Teacher.tname = ?", tname)

	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer rows.Close()
	var tid, cid, cname, classTime, credit string
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
func QueryCourseName(cid string) string {
	var cname string
	err := Db.QueryRow("select cname from Course where cid = ?", cid).Scan(&cname)
	if err != nil {
		log.Fatal(err)
	}
	return cname
}
func QueryCourseCredit(cid string) string {
	var credit string
	err := Db.QueryRow("select credit from Course where cid = ?", cid).Scan(&credit)
	if err != nil {
		log.Fatal(err)
	}
	return credit
}

func QueryTerm() []string {
	rows, err := Db.Query("select distinct(term) from CourseSchedule")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer rows.Close()
	term := make([]string, 0)
	var tmp string
	for rows.Next() {
		err = rows.Scan(&tmp)
		if err != nil {
			log.Fatal(err)
		}
		term = append(term, tmp)
	}
	return term
}

func QueryTeaName(tid string) string {
	var tname string
	err := Db.QueryRow("select tname from Teacher where tid = ?", tid).Scan(&tname)
	if err != nil {
		log.Fatal(err)
	}
	return tname
}
func QueryCourseWithTerm(term string) map[string][]string {
	ret := make(map[string][]string)
	rows, err := Db.Query("select Course.cid,cname,credit,tid,did,classTime from Course, CourseSchedule where Course.cid = CourseSchedule.cid and term = ?", term)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer rows.Close()
	var cid, tid, did, cname, classTime, credit string
	for rows.Next() {
		err := rows.Scan(&cid, &cname, &credit, &tid, &did, &classTime)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		ret["cid"] = append(ret["cid"], cid)
		ret["tid"] = append(ret["tid"], tid)
		ret["did"] = append(ret["did"], did)
		ret["tname"] = append(ret["tname"], QueryTeaName(tid))
		ret["dname"] = append(ret["dname"], QueryDepartmentName(did))
		ret["cname"] = append(ret["cname"], cname)
		ret["classTime"] = append(ret["classTime"], classTime)
		ret["credit"] = append(ret["credit"], credit)
	}
	return ret
}
