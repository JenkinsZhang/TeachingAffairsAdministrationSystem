package utils

func QueryCourseWithTermAndCid(term, cid string) (map[string][]string, error) {
	ret := make(map[string][]string)
	rows, err := Db.Query("select Teacher.tid,tname,classTime from CourseSchedule, Teacher where CourseSchedule.tid = Teacher.tid and term = ? and cid = ?", term, cid)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var tid, tname, classTime string
	for rows.Next() {
		err := rows.Scan(&tid, &tname, &classTime)
		if err != nil {
			return nil, err
		}
		ret["tid"] = append(ret["tid"], tid)
		ret["tname"] = append(ret["tname"], tname)
		ret["classTime"] = append(ret["classTime"], classTime)
	}
	return ret, nil
}
func QueryCourseNumberWithTerm(term string) (map[string][]interface{}, error) {
	ret := make(map[string][]interface{})
	rows, err := Db.Query("select cid, count(cid) from CourseSchedule where term = ? group by cid", term)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var cid string
	var cnt int
	ret["count"] = make([]interface{}, 0)
	for rows.Next() {
		err := rows.Scan(&cid, &cnt)
		if err != nil {
			return nil, err
		}
		ret["cid"] = append(ret["cid"], cid)
		ret["count"] = append(ret["count"], cnt)
	}
	return ret, nil
}
func GetCourse() (map[string][]string, error) {
	ret := make(map[string][]string)
	rows, err := Db.Query("select cid, cname, credit, did from Course")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var cid, cname, credit, did, dname string
	for rows.Next() {
		err := rows.Scan(&cid, &cname, &credit, &did)
		if err != nil {
			return nil, err
		}
		dname, err = QueryDepartmentName(did)
		if err != nil {
			return nil, err
		}
		ret["cid"] = append(ret["cid"], cid)
		ret["cname"] = append(ret["cname"], cname)
		ret["credit"] = append(ret["credit"], credit)
		ret["dname"] = append(ret["dname"], dname)
	}
	return ret, nil
}

func IfOpenSelectCourse() (string, error) {
	var msg string
	err := Db.QueryRow("select msg from Other where name = ?", "OpenCourseSelect").Scan(&msg)
	return msg, err
}

func QueryCourseWithCid(cid string) (map[string][]string, error) {
	term, err := GetCurrentTerm()
	if err != nil {
		return nil, err
	}
	ret := make(map[string][]string)
	rows, err := Db.Query("select Course.cid, cname, Teacher.tid, tname, credit, classTime from Course, CourseSchedule, Teacher where CourseSchedule.cid = Course.cid and CourseSchedule.tid = Teacher.tid and Course.cid = ? and term = ?", cid, term)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var tid, tname, cname, classTime, credit string
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

func QueryCourseWithCname(cname string) (map[string][]string, error) {
	term, err := GetCurrentTerm()
	if err != nil {
		return nil, err
	}
	ret := make(map[string][]string)
	rows, err := Db.Query("select Course.cid, cname, Teacher.tid, tname, credit, classTime from Course, CourseSchedule, Teacher where CourseSchedule.cid = Course.cid and CourseSchedule.tid = Teacher.tid and Course.cname = ? and term = ?", cname, term)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var cid, tid, tname, classTime, credit string
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

func QueryCourseWithTid(tid string) (map[string][]string, error) {
	term, err := GetCurrentTerm()
	if err != nil {
		return nil, err
	}
	ret := make(map[string][]string)
	rows, err := Db.Query("select Course.cid, cname, Teacher.tid, tname, credit, classTime from Course, CourseSchedule, Teacher where CourseSchedule.cid = Course.cid and CourseSchedule.tid = Teacher.tid and Teacher.tid = ? and term = ?", tid, term)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var cid, tname, cname, classTime, credit string
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
func QueryCourseWithTname(tname string) (map[string][]string, error) {
	term, err := GetCurrentTerm()
	if err != nil {
		return nil, err
	}
	ret := make(map[string][]string)
	rows, err := Db.Query("select Course.cid, cname, Teacher.tid, tname, credit, classTime from Course, CourseSchedule, Teacher where CourseSchedule.cid = Course.cid and CourseSchedule.tid = Teacher.tid and Teacher.tname = ? and term = ?", tname, term)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var tid, cid, cname, classTime, credit string
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
func QueryAllCourses() (map[string][]string, error) {
	term, err := GetCurrentTerm()
	if err != nil {
		return nil, err
	}
	ret := make(map[string][]string)
	rows, err := Db.Query("select Course.cid, cname, Teacher.tid, tname, credit, classTime from Course, CourseSchedule, Teacher where CourseSchedule.cid = Course.cid and CourseSchedule.tid = Teacher.tid and term = ?", term)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var tid, tname, cid, cname, classTime, credit string
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

func QueryDepartmentName(did string) (string, error) {
	var dname string
	err := Db.QueryRow("select dname from Department where did = ?", did).Scan(&dname)
	return dname, err
}
func QueryDepartmentId(dname string) (string, error) {
	var did string
	err := Db.QueryRow("select did from Department where dname = ?", dname).Scan(&did)
	return did, err
}

func QueryCourseName(cid string) (string, error) {
	var cname string
	err := Db.QueryRow("select cname from Course where cid = ?", cid).Scan(&cname)
	return cname, err
}
func QueryCourseCredit(cid string) (string, error) {
	var credit string
	err := Db.QueryRow("select credit from Course where cid = ?", cid).Scan(&credit)
	return credit, err
}

func QueryCourseWithTerm(term string) (map[string][]string, error) {
	ret := make(map[string][]string)
	rows, err := Db.Query("select Course.cid,cname,credit,tid,did,classTime from Course, CourseSchedule where Course.cid = CourseSchedule.cid and term = ?", term)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var cid, tid, did, cname, classTime, credit string
	for rows.Next() {
		err := rows.Scan(&cid, &cname, &credit, &tid, &did, &classTime)
		if err != nil {
			return nil, err
		}
		ret["cid"] = append(ret["cid"], cid)
		ret["tid"] = append(ret["tid"], tid)
		ret["did"] = append(ret["did"], did)
		tname, err := QueryTeaName(tid)
		if err != nil {
			return nil, err
		}
		ret["tname"] = append(ret["tname"], tname)
		dname, err := QueryDepartmentName(did)
		if err != nil {
			return nil, err
		}
		ret["dname"] = append(ret["dname"], dname)
		ret["cname"] = append(ret["cname"], cname)
		ret["classTime"] = append(ret["classTime"], classTime)
		ret["credit"] = append(ret["credit"], credit)
	}
	return ret, nil
}

func GetDepartment() (map[string][]string, error) {
	ret := make(map[string][]string)
	rows, err := Db.Query("select did,dname from Department")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var did, dname string
	for rows.Next() {
		err = rows.Scan(&did, &dname)
		if err != nil {
			return nil, err
		}
		ret["did"] = append(ret["did"], did)
		ret["dname"] = append(ret["dname"], dname)
	}
	return ret, nil
}

// -----------------------------------------
