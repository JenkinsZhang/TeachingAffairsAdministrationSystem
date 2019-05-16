package utils

import (
	"fmt"
)

// query student's courses in a specific term.
func QueryStuCourses(id, term string) (map[string][]string, error) {
	ret := make(map[string][]string)
	rows, err := Db.Query("select CourseCalendar.cid, cname,CourseCalendar.tid,tname, credit, classTime from CourseCalendar, Course, Teacher where CourseCalendar.id = ? and CourseCalendar.cid = Course.cid and CourseCalendar.tid = Teacher.tid and term = ?", id, term)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var cid, tid, tname, cname, classTime, credit string
	for rows.Next() {
		err := rows.Scan(&cid, &cname, &tid, &tname, &credit, &classTime)
		if err != nil {
			return nil, err
		}
		ret["cid"] = append(ret["cid"], cid)
		ret["tid"] = append(ret["tid"], tid)
		ret["tname"] = append(ret["tname"], tname)
		ret["cname"] = append(ret["cname"], cname)
		ret["classTime"] = append(ret["classTime"], classTime)
		ret["credit"] = append(ret["credit"], credit)
	}
	return ret, nil
}

func GetStudentProfile(id string) (map[string]string, error) {

	var name, gender, did, dname, phone, birthplace, birthday string
	err := Db.QueryRow("select name, gender, did,phone,birthplace,birthday from Student where id = ?", id).Scan(&name, &gender, &did, &phone, &birthplace, &birthday)
	if err != nil {
		return nil, err
	}
	dname, err = QueryDepartmentName(did)
	if err != nil {
		return nil, err
	}
	ret := map[string]string{
		"name":       name,
		"gender":     gender,
		"dname":      dname,
		"id":         id,
		"phone":      phone,
		"birthplace": birthplace,
		"birthday":   birthday,
	}
	return ret, nil
}

func GetAllStudentProfile() (map[string][]string, error) {
	ret := make(map[string][]string)
	rows, err := Db.Query("select id, name, gender, did,phone,birthplace,birthday from Student")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	info := [...]string{"id", "name", "gender", "did", "phone", "birthplace", "birthday"}
	val := make([]string, 10)
	for rows.Next() {
		err = rows.Scan(&val[0], &val[1], &val[2], &val[3], &val[4], &val[5], &val[6])
		if err != nil {
			return nil, err
		}
		for i, v := range info {
			ret[v] = append(ret[v], val[i])
		}
		dname, err := QueryDepartmentName(val[3])
		if err != nil {
			return nil, err
		}
		ret["dname"] = append(ret["dname"], dname)
	}
	return ret, nil
}

func QueryStuDepartmentId(id string) (string, error) {
	var did string
	err := Db.QueryRow("select did from Student where id = ? ", id).Scan(&did)
	return did, err
}

func QueryTotalStu(did string) (string, error) {
	var cnt int
	rows, err := Db.Query("select count(id) from Student where did = ?", did)
	defer rows.Close()
	if err != nil {
		return "", err
	}
	for rows.Next() {
		err := rows.Scan(&cnt)
		if err != nil {
			return "", err
		}
	}
	return fmt.Sprintf("%d", cnt), nil
}

func QueryStuName(id string) (string, error) {
	var name string
	err := Db.QueryRow("select name from Student where id = ?", id).Scan(&name)
	return name, err
}

// -----------------------------------------------------
