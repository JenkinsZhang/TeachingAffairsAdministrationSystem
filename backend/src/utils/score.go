package utils

import (
	"fmt"
	"math"
	"strconv"
)

func QueryStuCourseScore(id, term string) (map[string][]string, error) {
	ret := make(map[string][]string)
	rows, err := Db.Query("select CourseCalendar.cid, cname,CourseCalendar.tid,tname, credit, score from CourseCalendar, Course, Teacher where CourseCalendar.id = ? and CourseCalendar.cid = Course.cid and CourseCalendar.tid = Teacher.tid and term = ?", id, term)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var cid, tid, tname, cname, score, credit string
	for rows.Next() {
		err := rows.Scan(&cid, &cname, &tid, &tname, &credit, &score)
		if err != nil {
			return nil, err
		}
		if score == "-1" {
			score = ""
		}
		ret["cid"] = append(ret["cid"], cid)
		ret["tid"] = append(ret["tid"], tid)
		ret["tname"] = append(ret["tname"], tname)
		ret["cname"] = append(ret["cname"], cname)
		ret["score"] = append(ret["score"], score)
		ret["credit"] = append(ret["credit"], credit)
	}
	return ret, nil
}

func QueryStuAllCourseScore(id string) (map[string][]string, error) {
	cterm, _ := GetCurrentTerm()

	ret := make(map[string][]string)
	rows, err := Db.Query("select CourseCalendar.cid, cname,term, credit, score from CourseCalendar, Course where CourseCalendar.id = ? and CourseCalendar.cid = Course.cid and term != ?", id, cterm)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var cid, term, cname, score, credit string
	for rows.Next() {
		err := rows.Scan(&cid, &cname, &term, &credit, &score)
		if err != nil {
			return nil, err
		}
		ret["cid"] = append(ret["cid"], cid)
		ret["cname"] = append(ret["cname"], cname)
		ret["score"] = append(ret["score"], score)
		ret["credit"] = append(ret["credit"], credit)
		ret["term"] = append(ret["term"], term)
	}

	return ret, nil
}

func QueryTermAveScore(id, term string) (string, error) {
	fmt.Println(id, term)
	rows, err := Db.Query("select avg(score) from CourseCalendar where id = ? and term = ?", id, term)
	defer rows.Close()
	if err != nil {
		return "", err
	}
	var score float32
	for rows.Next() {
		err := rows.Scan(&score)
		fmt.Println(score)
		if err != nil {
			return "", err
		}
	}
	return fmt.Sprintf("%f", score), nil
}

func QueryStuAveScore(id string) (string, error) {
	cterm, _ := GetCurrentTerm()
	var score float64
	rows, err := Db.Query("select score,credit from CourseCalendar,Course where Course.cid = CourseCalendar.cid and id = ? and term != ?", id, cterm)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	sum := 0.0
	sumcre := 0.0
	var cre string
	for rows.Next() {
		err := rows.Scan(&score, &cre)
		if err != nil {
			return "", err
		}
		if score >= 0 {
			credit, _ := strconv.ParseFloat(cre, 64)
			sumcre += credit
			sum += score * credit
		}
	}
	eps := 1e-6
	if math.Abs(sumcre) < eps {
		return "0", nil
	}
	return fmt.Sprintf("%f", sum/sumcre), nil
}

func QueryStuRank(id string, did string) (string, error) {
	rows, err := Db.Query("select Student.id, avg(score) as s  from CourseCalendar, Student  where CourseCalendar.id = Student.id and  did = ? group by Student.id order by s desc", did)
	defer rows.Close()
	if err != nil {
		return "", err
	}

	cnt := 1
	var tmpid string
	var tmps float32

	for rows.Next() {
		err := rows.Scan(&tmpid, &tmps)
		if err != nil {
			return "", err
		}
		if tmpid == id {
			break
		}
		cnt++
	}
	return fmt.Sprintf("%d", cnt), nil

}
