package utils

func QueryTeaDepartmentId(tid string) (string, error) {
	var did string
	err := Db.QueryRow("select did from Teacher where tid = ? ", tid).Scan(&did)
	return did, err
}
func GetTeacherProfile(tid string) (map[string]string, error) {
	var tname, gender, did, dname string
	err := Db.QueryRow("select tname, gender, did from Teacher where tid = ?", tid).Scan(&tname, &gender, &did)
	if err != nil {
		return nil, err
	}
	dname, err = QueryDepartmentName(did)
	if err != nil {
		return nil, err
	}
	ret := map[string]string{
		"tname":  tname,
		"gender": gender,
		"dname":  dname,
		"tid":    tid,
	}
	return ret, nil
}

func QueryStuWithTidCid(tid, cid, term string) (map[string][]string, error) {
	var id, did, dname, score string
	rows, err := Db.Query("select id, score from CourseCalendar where tid = ? and cid = ? and term = ?", tid, cid, term)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	ret := make(map[string][]string)

	for rows.Next() {
		err := rows.Scan(&id, &score)
		if err != nil {
			return nil, err
		}
		did, err = QueryStuDepartmentId(id)
		if err != nil {
			return nil, err
		}
		dname, err = QueryDepartmentName(did)
		if err != nil {
			return nil, err
		}
		name, err := QueryStuName(id)
		if err != nil {
			return nil, err
		}
		ret["id"] = append(ret["id"], id)
		ret["name"] = append(ret["name"], name)
		ret["dname"] = append(ret["dname"], dname)
		ret["score"] = append(ret["score"], score)
	}
	return ret, nil
}

func QueryCourseScore(tid, cid, term string) ([]string, error) {
	rows, err := Db.Query("select score from CourseCalendar where tid = ? and cid = ? and term = ?", tid, cid, term)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	ret := make([]string, 0)
	var score string
	for rows.Next() {
		err := rows.Scan(&score)
		if err != nil {
			return nil, err
		}
		ret = append(ret, score)
	}
	return ret, nil
}

func QueryTeaCourse(tid, term string) (map[string][]string, error) {
	var cid, classTime, cname, credit string
	rows, err := Db.Query("select cid,classTime from CourseCalendar where tid = ? and term = ?", tid, term)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	ret := make(map[string][]string)

	for rows.Next() {
		err := rows.Scan(&cid, &classTime)
		if err != nil {
			return nil, err
		}
		cname, err = QueryCourseName(cid)
		if err != nil {
			return nil, err
		}
		credit, err = QueryCourseCredit(cid)
		if err != nil {
			return nil, err
		}
		ret["cid"] = append(ret["cid"], cid)
		ret["cname"] = append(ret["cname"], cname)
		ret["credit"] = append(ret["credit"], credit)
		ret["classTime"] = append(ret["classTime"], classTime)
	}
	return ret, nil
}

func QueryCourseStuInfo(tid, cid, term string) (map[string][]string, error) {
	var id string
	rows, err := Db.Query("select id from CourseCalendar where tid = ? and cid = ? and term = ?", tid, cid, term)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	ret := make(map[string][]string, 0)
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		c, err := GetStudentProfile(id)
		if err != nil {
			return nil, err
		}
		for key, val := range c {
			ret[key] = append(ret[key], val)
		}
	}
	return ret, nil
}

func QueryTeaName(tid string) (string, error) {
	var tname string
	err := Db.QueryRow("select tname from Teacher where tid = ?", tid).Scan(&tname)
	return tname, err
}

func GetAllTeacherProfile() (map[string][]string, error) {
	ret := make(map[string][]string)
	rows, err := Db.Query("select tid, tname, gender, did, wage,education,birthday from Teacher")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	info := [...]string{"tid", "tname", "gender", "did", "wage", "education", "birthday"}
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
