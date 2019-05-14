package utils

func QueryCourseWithCid(cid string) (map[string][]string, error) {
	ret := make(map[string][]string)
	rows, err := Db.Query("select Course.cid, cname, Teacher.tid, tname, credit, classTime from Course, CourseSchedule, Teacher where CourseSchedule.cid = Course.cid and CourseSchedule.tid = Teacher.tid and Course.cid = ?", cid)
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
	ret := make(map[string][]string)
	rows, err := Db.Query("select Course.cid, cname, Teacher.tid, tname, credit, classTime from Course, CourseSchedule, Teacher where CourseSchedule.cid = Course.cid and CourseSchedule.tid = Teacher.tid and Course.cname = ?", cname)
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
	ret := make(map[string][]string)
	rows, err := Db.Query("select Course.cid, cname, Teacher.tid, tname, credit, classTime,term from Course, CourseSchedule, Teacher where CourseSchedule.cid = Course.cid and CourseSchedule.tid = Teacher.tid and Teacher.tid = ?", tid)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	var cid, tname, cname, classTime, credit, term string
	for rows.Next() {
		err := rows.Scan(&cid, &cname, &tid, &tname, &credit, &classTime, &term)
		if err != nil {
			return nil, err
		}
		ret["cid"] = append(ret["cid"], cid)
		ret["tid"] = append(ret["tid"], tid)
		ret["tname"] = append(ret["tname"], tname)
		ret["cname"] = append(ret["cname"], cname)
		ret["classTime"] = append(ret["classTime"], classTime)
		ret["credit"] = append(ret["credit"], credit)
		ret["term"] = append(ret["term"], term)
	}
	return ret, nil
}
func QueryCourseWithTname(tname string) (map[string][]string, error) {
	ret := make(map[string][]string)
	rows, err := Db.Query("select Course.cid, cname, Teacher.tid, tname, credit, classTime from Course, CourseSchedule, Teacher where CourseSchedule.cid = Course.cid and CourseSchedule.tid = Teacher.tid and Teacher.tname = ?", tname)
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
	ret := make(map[string][]string)
	rows, err := Db.Query("select Course.cid, cname, Teacher.tid, tname, credit, classTime from Course, CourseSchedule, Teacher where CourseSchedule.cid = Course.cid and CourseSchedule.tid = Teacher.tid")
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

func QueryTerm() ([]string, error) {
	rows, err := Db.Query("select distinct(term) from CourseSchedule")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	term := make([]string, 0)
	var tmp string
	for rows.Next() {
		err = rows.Scan(&tmp)
		if err != nil {
			return nil, err
		}
		term = append(term, tmp)
	}
	return term, nil
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

// -----------------------------------------
