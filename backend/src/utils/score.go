package utils

import (
	"fmt"
	"log"
)

func QueryStuCourseScore(id, term string) map[string][]string {
	ret := make(map[string][]string)
	rows, err := Db.Query("select CourseCalendar.cid, cname,CourseCalendar.tid,tname, credit, score from CourseCalendar, Course, Teacher where CourseCalendar.id = ? and CourseCalendar.cid = Course.cid and CourseCalendar.tid = Teacher.tid and term = ?", id, term)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer rows.Close()
	var cid, tid, tname, cname, score, credit string
	for rows.Next() {
		err := rows.Scan(&cid, &cname, &tid, &tname, &credit, &score)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		ret["cid"] = append(ret["cid"], cid)
		ret["tid"] = append(ret["tid"], tid)
		ret["tname"] = append(ret["tname"], tname)
		ret["cname"] = append(ret["cname"], cname)
		ret["score"] = append(ret["score"], score)
		ret["credit"] = append(ret["credit"], credit)
	}
	return ret
}

func QueryStuAllCourseScore(id string) map[string][]string {
	ret := make(map[string][]string)
	rows, err := Db.Query("select CourseCalendar.cid, cname,CourseCalendar.tid,tname, credit, score from CourseCalendar, Course, Teacher where CourseCalendar.id = ? and CourseCalendar.cid = Course.cid and CourseCalendar.tid = Teacher.tid", id)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer rows.Close()
	var cid, tid, tname, cname, score, credit string
	for rows.Next() {
		err := rows.Scan(&cid, &cname, &tid, &tname, &credit, &score)
		if err != nil {
			log.Fatal(err)
			return nil
		}
		ret["cid"] = append(ret["cid"], cid)
		ret["tid"] = append(ret["tid"], tid)
		ret["tname"] = append(ret["tname"], tname)
		ret["cname"] = append(ret["cname"], cname)
		ret["score"] = append(ret["score"], score)
		ret["credit"] = append(ret["credit"], credit)
	}
	return ret
}

func QueryTermAveScore(id, term string) string {
	rows, err := Db.Query("select avg(score) from CourseCalendar where id = ? and term = ?", id, term)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var score float32
	for rows.Next() {
		err := rows.Scan(&score)
		if err != nil {
			log.Fatal(err)
		}
	}
	return fmt.Sprintf("%f", score)
}

func QueryStuAveScore(id string) string {
	rows, err := Db.Query("select avg(score) from CourseCalendar where id = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var score float32
	for rows.Next() {
		err := rows.Scan(&score)
		if err != nil {
			log.Fatal(err)
		}
	}
	return fmt.Sprintf("%f", score)
}

func QueryStuRank(id string, did string) string {
	rows, err := Db.Query("select Student.id, avg(score) as s  from CourseCalendar, Student  where CourseCalendar.id = Student.id and  did = ? group by Student.id order by s desc", did)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	cnt := 1
	var tmpid string
	var tmps float32

	for rows.Next() {
		err := rows.Scan(&tmpid, &tmps)
		if err != nil {
			log.Fatal(err)
		}
		if tmpid == id {
			break
		}
		cnt++
	}
	return fmt.Sprintf("%d", cnt)

}
